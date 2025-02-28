proto:
	rm -f common/pb/*.go
	protoc --proto_path common/proto --go_out common/pb --go_opt=paths=source_relative \
	--go-grpc_out common/pb --go-grpc_opt=paths=source_relative \
 common/proto/*.proto

PHONY: proto