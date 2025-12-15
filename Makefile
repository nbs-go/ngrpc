NGRPC_DIR := $(shell pwd)
NGRPC_PROTO_SRC_DIR := ${NGRPC_DIR}/proto
NGRPC_PROTO_SRC_FILES := $(shell find ${NGRPC_PROTO_SRC_DIR} -type f -name "*.proto")
NGRPC_PROTO_GO_OUT_DIR := ${NGRPC_DIR}

## proto: Generate Go files for package ngrpc
.PHONY: proto
proto:
	@-echo "  > proto: Removing generated Go files..."
	@-rm ${NGRPC_PROTO_GO_OUT_DIR}/*.pb.go
	@-echo "  > proto: Generate Go files from proto..."
	@protoc --proto_path=${NGRPC_PROTO_SRC_DIR} \
		--go_out=${NGRPC_PROTO_GO_OUT_DIR} \
		--go_opt paths=source_relative \
		--go-grpc_out=${NGRPC_PROTO_GO_OUT_DIR} \
		--go-grpc_opt paths=source_relative \
		${NGRPC_PROTO_SRC_DIR}/*.proto
	@-echo "  > proto: Done"

## configure: Install toolchain

GOBIN = ${NGRPC_DIR}/.tmp/go/bin
PROTOC_GEN_GO_VERSION := "v1.36.10"
PROTOC_GEN_GO_GRPC_VERSION := "v1.6.0"
PROTOBUF_VERSION := "31.1"

.PHONY: configure
configure:
	@-echo "  > configure: Installing generators for Go gRPC..."
	@-(mise ls -i protobuf | grep ${PROTOBUF_VERSION}) || MISE_HTTP_TIMEOUT=60 mise use protobuf@${PROTOBUF_VERSION}
	@-go install google.golang.org/protobuf/cmd/protoc-gen-go@${PROTOC_GEN_GO_VERSION}
	@-echo "  > configure: Done"