# gRPC Project 

## Dependences 

Install [protocol buffer compiler](https://grpc.io/docs/protoc-installation/).

## How to start for Go.

First, install the depencences. 
```bash 
$ go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
$ go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
$ go get google.golang.org/protobuf
$ go get github.com/lib/pq 
$ go get google.golang.org/grpc
```

Secondary, make a student.proto file, we transform the file to a Go language using this command. 
```bash 
$ protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative studentpb/student.proto
```

## Why protobuffer? 

Json, if Go, Python, PHP, Java need to read this json; all the leanguages need to serializated and deserializated, these processes take more time than a protobuffers.

```json 
{
    "foo": "bar"
}
```
Json is easy to
+ read
+ flexible
- serialization
- deserialization
+ web standard

Protobuffers have a x.proto and it is agnostic to the language program and is read one time on the start.

In comparision:

| Json | ProtoBuffers |
|---|---|
| + read | - read |
| + flexible | - flexible |
| - serialization/deserialization | + serialization/deserialization|
| + web standard | + reutilization |