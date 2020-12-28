# grpc-go-example
A simple gRPC client and server in golang to illustrate how gRPC works.


## What does it do
Looking at the `hash.proto` file we can see 3 routes:
```protobuf
service HashService {
    rpc Sha1Message(Message) returns (Message) {}
    rpc Sha265Message(Message) returns (Message) {}
    rpc Sha512Message(Message) returns (Message) {}
}
```
Each one hashing the incoming Message according to it's name (SHA1, 265, 512).

The message declaration looks like this:
```protobuf
message Message {
    string body = 1;
}
```

## Generate code
To generate the `hash/hash.pb.go` file run:

```
protoc --go_out=plugins=grpc:hash hash.proto
```

## Run it
```
go run server.go
```
```
go run client.go
```
