package main

import (
  "context"
  pb "github.com/ckbball/smurfin-catalog/proto/catalog"
  "go.mongodb.org/mongo-driver/bson"
  "go.mongodb.org/mongo-driver/mongo"
  "go.mongodb.org/mongo-driver/mongo/options"
)

type repository interface {
  FindItems(spec *pb.Specification) ([]*pb.Item, error)
  CreateItem(item *pb.Item) error
  RemoveItem(item *pb.Item) error
  ListItems(spec *pb.ListSpecification) ([]*pb.Item, error)
}

type ItemRepository struct {
  collection *mongo.Collection
}

func (repository *ItemRepository) FindItems(spec *pb.Specification) ([]*pb.Item, error) {
  // make bson filter object here
  filter := bson.D{
    {
      "Solo",
      &bson.A{
        "$eq",
        spec.Solo,
      },
    },
    {
      "PriceDollar",
      &bson.A{
        "$lte",
        spec.PriceDollars,
      },
    },
    {
      "VendorId",
      &bson.A{
        "$eq",
        spec.VendorId,
      },
    },
  }

  findOptions := options.Find()
  findOptions.SetLimit(20)
  findOptions.SetSort(bson.D{{"_id", -1}})
  findOptions.SetSkip(int64(spec.PageNum))

  var items []*pb.Item
  cur, err := repository.collection.Find(context.TODO(), filter, findOptions)
  if err != nil {
    return nil, err
  }
  defer cur.Close(context.TODO())

  for cur.Next(context.TODO()) {
    var elem *pb.Item
    err := cur.Decode(&elem)
    if err != nil {
      return nil, err
    }

    items = append(items, elem)
  }

  if err := cur.Err(); err != nil {
    return items, err
  }

  return items, nil
}

func (repository *ItemRepository) CreateItem(item *pb.Item) error {
  _, err := repository.collection.InsertOne(context.TODO(), item)
  return err
}

func (repository *ItemRepository) RemoveItem(item *pb.Item) error {
  filter := bson.D{{"_id", item.Id}}

  // change _ to var to capture number of items deleted
  _, err := repository.collection.DeleteOne(context.TODO(), filter)
  if err != nil {
    return err
  }
  return err
}

func (repository *ItemRepository) ListItems(spec *pb.ListSpecification) ([]*pb.Item, error) {

  findOptions := options.Find()
  findOptions.SetLimit(20)
  findOptions.SetSort(bson.D{{"_id", -1}})
  findOptions.SetSkip(int64(spec.Page))

  var items []*pb.Item
  cur, err := repository.collection.Find(context.TODO(), bson.D{{}}, findOptions)
  if err != nil {
    return nil, err
  }
  defer cur.Close(context.TODO())

  for cur.Next(context.TODO()) {
    var elem *pb.Item
    err := cur.Decode(&elem)
    if err != nil {
      return nil, err
    }

    items = append(items, elem)
  }

  if err := cur.Err(); err != nil {
    return items, err
  }

  return items, nil
}
