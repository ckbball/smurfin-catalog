package v1

import (
  "context"
  "database/sql"
  "fmt"
  //"time"

  //"github.com/golang/protobuf/ptypes"
  "google.golang.org/grpc/codes"
  "google.golang.org/grpc/status"

  v1 "github.com/ckbball/smurfin-catalog/pkg/api/v1"
)

const (
  apiVersion = "v1"
)

type catalogServiceServer struct {
  db *sql.DB
}

func NewCatalogServiceServer(db *sql.DB) v1.CatalogServiceServer {
  return &catalogServiceServer{db: db}
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

  // Query database
  rows, err := c.QueryContext(ctx, "SELECT * FROM items WHERE id > ? AND solo=? OR flex=? ORDER BY id ASC LIMIT 20", item_id, req.Solo, req.Flex)
  if err != nil {
    return nil, status.Error(codes.Unknown, "failed to query items: "+err.Error())
  }
  defer rows.Close()

  list := []*v1.Item{}
  for rows.Next() {
    it := new(v1.Item)
    if err := rows.Scan(&it.Id, &it.VendorId, &it.BlueEssence, &it.RiotPoints, &it.Solo, &it.Flex, &it.PriceDollars, &it.PriceCents, &it.Level, &it.Email, &it.Password, &it.Login, &it.LoginPassword); err != nil {
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
    req.Item.VendorId, req.Item.BlueEssence, req.Item.RiotPoints, req.Item.Solo, req.Item.Flex, req.Item.PriceDollars, req.Item.PriceCents, req.Item.Level, req.Item.Email, req.Item.Password, req.Item.Login, req.Item.LoginPassword)
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

  // Query database
  rows, err := c.QueryContext(ctx, "SELECT * FROM items WHERE id > ? ORDER BY id ASC LIMIT 20", item_id)
  if err != nil {
    return nil, status.Error(codes.Unknown, "failed to query items: "+err.Error())
  }
  defer rows.Close()

  list := []*v1.Item{}
  for rows.Next() {
    it := new(v1.Item)
    if err := rows.Scan(&it.Id, &it.VendorId, &it.BlueEssence, &it.RiotPoints, &it.Solo, &it.Flex, &it.PriceDollars, &it.PriceCents, &it.Level, &it.Email, &it.Password, &it.Login, &it.LoginPassword); err != nil {
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

}
