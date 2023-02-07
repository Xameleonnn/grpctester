all: ech set_PATH generate

ech:
	go install google.golang.org/protobuf/cmd/protoc-gen-go \
	go install google.golang.org/grpc/cmd/protoc-gen-go-grpc

set_PATH:
	export PATH=$PATH:$(pwd)

generate:
	protoc -I=. --go_out=. --go-grpc_out=. grpc.proto
