cd cmd/server
go build .
.\server.exe -grpc-port=9090 -http-port=8080 -db-host=blah -db-user=dev -db-password=dev-user5 -db-schema=catalog -redis-address=6379
