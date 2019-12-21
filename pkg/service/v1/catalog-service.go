package v1

import (
  "context"
  "database/sql"
  "fmt"
  "strconv"
  //"time"

  //"github.com/golang/protobuf/ptypes"
  "github.com/go-redis/cache/v7"
  "google.golang.org/grpc/codes"
  "google.golang.org/grpc/status"

  v1 "github.com/ckbball/smurfin-catalog/pkg/api/v1"
)

const (
  apiVersion = "v1"
)

type catalogServiceServer struct {
  db     *sql.DB
  rcache *cache.Codec
}

func NewCatalogServiceServer(db *sql.DB, rcache *cache.Codec) v1.CatalogServiceServer {
  return &catalogServiceServer{db: db, rcache: rcache}
}

func (s *catalogServiceServer) checkAPI(api string) error {
  if len(api) > 0 {
    if apiVersion != api {
      return status.Errorf(codes.Unimplemented,
        "unsupported API version: service implements API version '%s', but asked for '%s'", apiVersion, api)
    }
  }
  return nil
}

func (s *catalogServiceServer) connect(ctx context.Context) (*sql.Conn, error) {
  c, err := s.db.Conn(ctx)
  if err != nil {
    return nil, status.Error(codes.Unknown, "failed to connect to database-> "+err.Error())
  }
  return c, nil
}

// Other service funcs

// handler for getbyid endpoint
// takes an id in the url path
// returns item matching id
func (s *catalogServiceServer) GetById(ctx context.Context, req *v1.GetByIdRequest) (*v1.GetByIdResponse, error) {
  // check api version
  if err := s.checkAPI(req.Api); err != nil {
    return nil, err
  }

  // access query params through fields in message of proto files.

  //check cache
  item := new(v1.Item)
  err := s.rcache.Get(strconv.Itoa(int(req.Id)), &item)
  if err == nil && item != nil {
    // return proper stuff
    return &v1.GetByIdResponse{
      Api:  apiVersion,
      Item: item,
    }, nil
  }

  // get SQL connection
  c, err := s.connect(ctx)
  if err != nil {
    return nil, err
  }
  defer c.Close()

  // db get
  rows, err := c.QueryContext(ctx, "SELECT id, vendorid, blueessence, riotpoints, solo, flex, pricedollars, pricecents, level FROM items WHERE id=?", req.Id)
  if err != nil {
    return nil, status.Error(codes.Unknown, "failed to get item by id-> "+err.Error())
  }
  defer rows.Close()

  if !rows.Next() {
    if err := rows.Err(); err != nil {
      return nil, status.Error(codes.Unknown, "failed to retrieve data from items-> "+err.Error())
    }
    return nil, status.Error(codes.NotFound, fmt.Sprintf("items with ID='%d' is not found",
      req.Id))
  }

  if err = rows.Scan(&item.Id, &item.VendorId, &item.BlueEssence, &item.RiotPoints, &item.Solo, &item.Flex, &item.PriceDollars, &item.PriceCents, &item.Level); err != nil {
    return nil, status.Error(codes.Unknown, "failed to retrieve field values from ToDo row-> "+err.Error())
  }
  if rows.Next() {
    return nil, status.Error(codes.Unknown, fmt.Sprintf("found multiple items rows with ID='%d'",
      req.Id))
  }
  return &v1.GetByIdResponse{
    Api:  apiVersion,
    Item: item,
  }, nil
}

// handler for find items endpoint
// takes a list of fields to filter the query for
// returns a list of items that match given query
func (s *catalogServiceServer) FindItems(ctx context.Context, req *v1.Specification) (*v1.Response, error) {
  // check api version
  if err := s.checkAPI(req.Api); err != nil {
    return nil, err
  }

  // get SQL connection
  c, err := s.connect(ctx)
  if err != nil {
    return nil, err
  }
  defer c.Close()

  item_id := 0
  if req.PageNum > 1 {
    item_id = 20 * int((req.PageNum - 1))
  }

  // add caching logic here

  // Query database
  rows, err := c.QueryContext(ctx, "SELECT id, vendorid, blueessence, riotpoints, solo, flex, pricedollars, pricecents, level FROM items WHERE id > ? AND solo=? OR flex=? ORDER BY id ASC LIMIT 20", item_id, req.Solo, req.Flex)
  if err != nil {
    return nil, status.Error(codes.Unknown, "failed to query items: "+err.Error())
  }
  defer rows.Close()

  list := []*v1.Item{}
  for rows.Next() {
    it := new(v1.Item)
    if err := rows.Scan(&it.Id, &it.VendorId, &it.BlueEssence, &it.RiotPoints, &it.Solo, &it.Flex, &it.PriceDollars, &it.PriceCents, &it.Level); err != nil {
      return nil, status.Error(codes.Unknown, "failed to retrieve field values from item row-> "+err.Error())
    }
    list = append(list, it)
  }

  if err := rows.Err(); err != nil {
    return nil, status.Error(codes.Unknown, "failed to retrieve data from item-> "+err.Error())
  }

  return &v1.Response{
    Api:   apiVersion,
    Items: list,
  }, nil

}

func (s *catalogServiceServer) Create(ctx context.Context, req *v1.CreateRequest) (*v1.CreateResponse, error) {
  // check api version
  if err := s.checkAPI(req.Api); err != nil {
    return nil, err
  }

  // get SQL connection
  c, err := s.connect(ctx)
  if err != nil {
    return nil, err
  }
  defer c.Close()

  res, err := c.ExecContext(ctx, `INSERT INTO items (VendorId, BlueEssence, RiotPoints, Solo, Flex, PriceDollars, PriceCents, Level, Email, Password, Login, LoginPassword) VALUES(?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`,
    req.Item.VendorId, req.Item.BlueEssence, req.Item.RiotPoints, req.Item.Solo, req.Item.Flex, req.Item.PriceDollars, req.Item.PriceCents, req.Item.Level, req.Item.Email, req.Item.EmailPassword, req.Item.LoginName, req.Item.LoginPassword)
  if err != nil {
    return nil, status.Error(codes.Unknown, "failed to insert into item-> "+err.Error())
  }
  id, err := res.LastInsertId()
  if err != nil {
    return nil, status.Error(codes.Unknown, "failed to retrieve id for created item-> "+err.Error())
  }

  return &v1.CreateResponse{
    Api: apiVersion,
    Id:  id,
  }, nil
}

// handles requests for the removeitem endpoint
// takes an item id and api version
// returns number of items deleted and api version
func (s *catalogServiceServer) RemoveItem(ctx context.Context, req *v1.RemoveRequest) (*v1.RemoveResponse, error) {
  // check api version
  if err := s.checkAPI(req.Api); err != nil {
    return nil, err
  }

  // get SQL connection
  c, err := s.connect(ctx)
  if err != nil {
    return nil, err
  }
  defer c.Close()

  // db operation to delete specified item
  res, err := c.ExecContext(ctx, "DELETE FROM items WHERE ID=?", req.Id)
  if err != nil {
    return nil, status.Error(codes.Unknown, "failed to delete item-> "+err.Error())
  }

  rows, err := res.RowsAffected()
  if err != nil {
    return nil, status.Error(codes.Unknown, "failed to get rows affected-> "+err.Error())
  }
  if rows == 0 {
    return nil, status.Error(codes.Unknown, fmt.Sprintf("item with ID ='%d' not found", req.Id))
  }

  return &v1.RemoveResponse{
    Api:     apiVersion,
    Deleted: rows,
  }, nil
}

func (s *catalogServiceServer) ListItems(ctx context.Context, req *v1.ListRequest) (*v1.ListResponse, error) {
  // check api version
  if err := s.checkAPI(req.Api); err != nil {
    return nil, err
  }

  // get SQL connection
  c, err := s.connect(ctx)
  if err != nil {
    return nil, err
  }
  defer c.Close()

  item_id := 0
  if req.Page > 1 {
    item_id = 20 * int((req.Page - 1))
  }

  // add caching logic here

  // Query database
  rows, err := c.QueryContext(ctx, "SELECT id, vendorid, blueessence, riotpoints, solo, flex, pricedollars, pricecents, level FROM items WHERE id > ? ORDER BY id ASC LIMIT 20", item_id)
  if err != nil {
    return nil, status.Error(codes.Unknown, "failed to query items: "+err.Error())
  }
  defer rows.Close()

  list := []*v1.Item{}
  for rows.Next() {
    it := new(v1.Item)
    if err := rows.Scan(&it.Id, &it.VendorId, &it.BlueEssence, &it.RiotPoints, &it.Solo, &it.Flex, &it.PriceDollars, &it.PriceCents, &it.Level); err != nil {
      return nil, status.Error(codes.Unknown, "failed to retrieve field values from item row-> "+err.Error())
    }
    list = append(list, it)
  }

  if err := rows.Err(); err != nil {
    return nil, status.Error(codes.Unknown, "failed to retrieve data from item-> "+err.Error())
  }

  return &v1.ListResponse{
    Api:   apiVersion,
    Items: list,
  }, nil

  // will need to add a function to get private info for the email to send
}
