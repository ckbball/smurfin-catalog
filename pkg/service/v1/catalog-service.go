package v1

import (
  "context"
  "database/sql"
  "fmt"
  "time"

  "github.com/golang/protobuf/ptypes"
  "google.golang.org/grpc/codes"
  "google.golang.org/grpc/status"

  "github.com/ckbball/smurfin-catalogue/pkg/api/v1"
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
