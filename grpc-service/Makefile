generate-protobuf:
	@protoc -I=proto --go_out=paths=source_relative:grpc_server \
	--go-grpc_out=paths=source_relative:grpc_server proto/*.proto