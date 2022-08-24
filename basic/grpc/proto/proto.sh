#/bin/bash
protoc --go_out=./gen/go --go_opt=paths=source_relative --go-grpc_out=./gen/go --go-grpc_opt=paths=source_relative --grpc-gateway_out=./gen/go --grpc-gateway_opt paths=source_relative trip.proto

protoc --grpc-gateway_out=paths=source_relative,grpc_api_configuration=trip.yaml:gen/go trip.proto