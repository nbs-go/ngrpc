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