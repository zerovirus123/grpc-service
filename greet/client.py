import grpc
from grpc_client.greet_pb2_grpc import GreetServiceStub
from grpc_client.greet_pb2 import GreetingRequest

def main():
    
    with grpc.insecure_channel("localhost:8080") as channel:
        client = GreetServiceStub(channel)
        request = GreetingRequest(name="Brian")
        response = client.Greeting(request)
        print(response)

main()