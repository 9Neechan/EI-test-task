proto:
	rm -f api/pb/*.go
	protoc --proto_path api/proto --go_out api/pb --go_opt=paths=source_relative \
	--go-grpc_out api/pb --go-grpc_opt=paths=source_relative \
 api/proto/*.proto

PHONY: proto