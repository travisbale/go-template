# Protocol Buffers

This directory contains Protocol Buffer (.proto) definitions for gRPC services.

## Structure

```
proto/
├── service.proto         # Your gRPC service definitions
├── Dockerfile           # Docker image for protoc compilation
└── README.md
```

## Generating Code

To generate Go code from .proto files:

```bash
make protoc
```

This will:
1. Build a Docker image with protoc and required plugins
2. Generate Go code in `internal/pb/`
3. Generate both protobuf messages (`.pb.go`) and gRPC service stubs (`_grpc.pb.go`)

## Adding New Services

1. Create a new `.proto` file or add to existing one:

```protobuf
syntax = "proto3";

package myservice;

option go_package = "github.com/yourorg/myservice/internal/pb";

service MyService {
  rpc GetResource(GetResourceRequest) returns (GetResourceResponse);
}

message GetResourceRequest {
  string id = 1;
}

message GetResourceResponse {
  string id = 1;
  string name = 2;
}
```

2. Run `make protoc` to generate Go code

3. Implement the service in `internal/api/grpc/`

4. Register the service in `internal/api/grpc/server.go`

## Proto3 Resources

- [Protocol Buffers Language Guide](https://protobuf.dev/programming-guides/proto3/)
- [gRPC Go Quick Start](https://grpc.io/docs/languages/go/quickstart/)
- [Style Guide](https://protobuf.dev/programming-guides/style/)
