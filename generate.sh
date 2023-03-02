rm -rf pb/*.go
protoc --proto_path=proto --go_out=pb --go_out=paths=source_relative \
  --go-grpc_out=pb --go-grpc_opt=paths=source_relative \
  proto/*.proto