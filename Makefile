API_PATH = docs/pb/v1
PROTO_API_DIR = docs/pb/v1
PROTO_OUT_DIR = pkg/papi
TOOL_IMG = "local.go.tool:latest"

proto-img-build:
	docker build -t ${TOOL_IMG} -f docker/tools/Dockerfile .

proto:
	docker run -it --rm --name gotool -v "${PWD}:/app" ${TOOL_IMG} make proto-build

proto-build:
	rm -rf ${PROTO_OUT_DIR}
	mkdir -p ${PROTO_OUT_DIR}
	protoc \
		-I ${API_PATH} \
		-I${GOPATH}/pkg/mod/github.com/grpc-ecosystem/grpc-gateway@v1.16.0/third_party/googleapis \
		-I /var/proto/protoc3/include \
		--include_imports \
		--go_out=$(PROTO_OUT_DIR) --go_opt=paths=source_relative \
        --go-grpc_out=$(PROTO_OUT_DIR)  --go-grpc_opt=paths=source_relative \
		--descriptor_set_out=$(PROTO_OUT_DIR)/api.pb \
		--grpc-gateway_out=logtostderr=true:./pkg/ \
		--swagger_out=allow_merge=true,merge_file_name=docs/api/api:. \
		./${PROTO_API_DIR}/*.proto
