# Setting Up Client Environment
Run `pip install -r requirements.txt`

Activate venv with `source venv/bin/activate`

# Converting Proto to gRPC Code

To convert greet.proto interfaces to gRPC code, run 

`python -m grpc_tools.protoc -I ../grpc-client/greet/proto --python_out=. --grpc_python_out=. ../grpc-client/greet/proto/greet.proto`