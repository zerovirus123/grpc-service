python -m grpc_tools.protoc -I ../grpc-client/greet/proto --python_out=. \
         --grpc_python_out=. ../grpc-client/greet/proto/greet.proto