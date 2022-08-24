#how to use


## 安装
```
安装protobuf
https://github.com/protocolbuffers/protobuf

https://github.com/grpc-ecosystem/grpc-gateway/
go install \
github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway \
github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2 \
google.golang.org/protobuf/cmd/protoc-gen-go \
google.golang.org/grpc/cmd/protoc-gen-go-grpc

```

## 运行
```
执行 sh proto/proto.sh 会生成gen/go/目录下3个文件
go run server/main.go
go run client/main.go
```

## http返回enum to int
https://grpc-ecosystem.github.io/grpc-gateway/docs/development/grpc-gateway_v2_migration_guide/

## note
```
struct=>byte=>json
```
