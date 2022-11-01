# GRPC Service
A simple demonstration of a gRPC protocol, using a Python client and a Golang server.

## Instructions

### Setting up the Client Environment
Create a virtual environment with `python -m venv venv`

Activate the virtual environment with `source venv/bin/activate`

Install the dependencies `pip install -r requirements.txt`

### Convert Proto to Client gRPC Code
To convert greet.proto interfaces to gRPC code, run `python -m grpc_tools.protoc -I proto --python_out=. --grpc_python_out=. proto/greet.proto`

### Convert Proto to Server gRPC Code

First install the protocol buffer compiler.

For Linux users, run `apt install -y protobuf-compiler`.

For Mac users, run `brew install protobuf`

Windows user can find the binaries at the [release page](https://github.com/protocolbuffers/protobuf/releases/tag/v3.17.3).

Install Go plugins for the protocol buffer compiler `go get -u github.com/golang/protobuf/protoc-gen-go`

In the Makefile's directory, run `make generate-protobuf`

# Run Server
Make sure that your system has Go working.

Run the server with `go run .`