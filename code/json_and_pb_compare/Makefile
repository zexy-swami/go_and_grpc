.PHONY: generate build run clean
.DEFAULT_GOAL := build

generate:
	mkdir -p pkg/pb
	protoc --go_out=pkg/pb/ api/proto/user.proto

build: generate
	CGO_ENABLED=0 GOOS=darwin GOARCH=arm64 \
	go build -o ./bin/json_and_pb_compare -ldflags "-s -w" -trimpath ./cmd/json_and_pb_compare/.

run:
	./bin/json_and_pb_compare

clean:
	rm -rf ./bin
	rm -rf ./pkg/pb