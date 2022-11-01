generate-protobuf:
	@protoc -I=proto --go_out=paths=source_relative:proto_out \
	--go-grpc_out=paths=source_relative:proto_out proto/*.proto