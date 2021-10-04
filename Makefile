create:
	protoc -I . \
		--go_out ./gen/go/pb/ --go_opt paths=source_relative \
		--go-grpc_out ./gen/go/pb/ --go-grpc_opt paths=source_relative \
		fibonacci.proto
	protoc -I . --grpc-gateway_out ./gen/go/pb/ \
		--grpc-gateway_opt logtostderr=true \
		--grpc-gateway_opt paths=source_relative \
		fibonacci.proto
build:
	go build -v ./cmd/server
	go build -v ./cmd/client
run: build
	./server
