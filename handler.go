package main

import (
  "context"
  pb "github.com/ckbball/smurfin-catalog/proto/catalog"
)

type handler struct {
  repo repository
}

// FindItems - finds items based on a search specification
// takes a proto.specification as input
// returns - a list of item(s)
func (s *handler) FindItems(ctx context.Context, req *pb.Specification, res *pb.Response) error {
  // Find the list of items that match spec
  items, err := s.repo.FindItems(req)
  if err != nil {
    return err
  }

  // Set the items as part of the response message
  res.Items = items
  return nil
}

// Create a new Item
func (s *handler) Create(ctx context.Context, req *pb.Item, res *pb.Response) error {
  if err := s.repo.CreateItem(req); err != nil {
    return err
  }
  res.Item = req
  return nil
}

// Remove an Item
func (s *handler) RemoveItem(ctx context.Context, req *pb.Specification, res *pb.Response) error {
  if err := s.repo.RemoveItem(req.Item); err != nil {
    return err
  }
  res.Item = req.Item
  return nil
}

// List items, paginated ?
func (s *handler) ListItems(ctx context.Context, req *pb.ListSpecification, res *pb.Response) error {
  items, err := s.repo.ListItems(req)
  if err != nil {
    return err
  }

  res.Items = items
  return nil
}
