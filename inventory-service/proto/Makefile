# gen_proto:
# 	protoc -I=. --go_out=. --go_opt=paths=source_relative inventory.proto
# 	protoc -I=. --go-grpc_out=. --go-grpc_opt=paths=source_relative inventory.proto
# 	protoc -I=. --grpc-gateway_out=. --grpc-gateway_opt=grpc_api_configuration=inventory.yaml --grpc-gateway_opt=paths=source_relative inventory.proto
# 	protoc -I=. --swagger_out=. --swagger_opt=grpc_api_configuration=inventory.yaml inventory.proto

all: proto gateway swagger

proto:
	protoc -I=. \
	  --go_out=. --go_opt=paths=source_relative \
	  --go-grpc_out=. --go-grpc_opt=paths=source_relative \
	  inventory.proto

gateway:
	protoc -I=. \
	  --grpc-gateway_out=. --grpc-gateway_opt=paths=source_relative \
	  --grpc-gateway_opt=grpc_api_configuration=inventory.yaml \
	  inventory.proto

swagger:
	protoc -I=. \
	  --swagger_out=. --swagger_opt=grpc_api_configuration=inventory.yaml \
	  inventory.proto

install-tools:
	go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
	go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
	go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway@latest
	go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-swagger@latest