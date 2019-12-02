package main

import (
  "context"
  "fmt"
  pb "github.com/ckbball/smurfin-catalog/proto/catalog"
  "github.com/micro/go-micro"
  "log"
  "os"
)

const (
  defaultHost = "datastore:27017"
)

func main() {
  srv := micro.NewService(
    micro.Name("go.micro.api.catalog"),
  )

  srv.Init()

  uri := os.Getenv("DB_HOST")
  if uri == "" {
    uri = defaultHost
  }

  client, err := CreateClient(uri)
  if err != nil {
    log.Panic(err)
  }
  defer client.Disconnect(context.TODO())

  itemCollection := client.Database("smurfin").Collection("item")
  repository := &ItemRepository{
    itemCollection,
  }

  pb.RegisterCatalogServiceHandler(srv.Server(), &handler{repository})

  if err := srv.Run(); err != nil {
    fmt.Println(err)
  }
}
