.PHONY: generate build_server build_client run_server run_client clean
.DEFAULT_GOAL := build_server

generate:
	mkdir -p pkg/pb
	protoc --go_out=pkg/pb/ --go-grpc_out=pkg/pb/ \
	api/proto/auth_service.proto

build_server:
	CGO_ENABLED=0 GOOS=darwin GOARCH=arm64 \
	go build -o ./bin/grpc_server -ldflags "-s -w" -trimpath ./cmd/grpc_server/.

build_client:
	CGO_ENABLED=0 GOOS=darwin GOARCH=arm64 \
	go build -o ./bin/grpc_client -ldflags "-s -w" -trimpath ./cmd/grpc_client/.

run_server:
	./bin/grpc_server

run_client:
	./bin/grpc_client

clean:
	rm -rf ./bin
	rm -rf ./pkg/pb
