# Convert Proto to gRPC code
First install the protocol buffer compiler.

For Linux users, run `apt install -y protobuf-compiler`.

For Mac users, run `brew install protobuf`

Windows user can find the binaries at the [release page](https://github.com/protocolbuffers/protobuf/releases/tag/v3.17.3).

Install Go plugins for the protocol buffer compiler `go get -u github.com/golang/protobuf/protoc-gen-go`

In the Makefile's directory, run `make generate-protobuf name=greet`

# Run Server
Make sure that your system has Go working.
Run the server with `go run .`