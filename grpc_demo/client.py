'''
    grpc python 客户端 demo
'''
import grpc
from hellogrpc import hellogrpc_pb2, hellogrpc_pb2_grpc

if __name__ == '__main__':
    with grpc.insecure_channel("localhost:50051") as channel:
        stub = hellogrpc_pb2_grpc.GreeterStub(channel)
        hello_request = hellogrpc_pb2.HelloRequest()
        hello_request.name = "abc"
        response: hellogrpc_pb2.HelloReply = stub.SayHello(hello_request)
        print(response.message)