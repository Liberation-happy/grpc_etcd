import grpc
from concurrent import futures

from proto import helloworld_pb2, helloworld_pb2_grpc


class Greeter(helloworld_pb2_grpc.GreeterServicer):
    def SayHello(self, request, context):
        print(request.name)
        return helloworld_pb2.HelloResponse(message=f"你好,{request.name}")


if __name__ == '__main__':
    # 1. 实例化server
    server = grpc.server(futures.ThreadPoolExecutor(max_workers=10))
    # 2. 注册逻辑到server
    helloworld_pb2_grpc.add_GreeterServicer_to_server(Greeter(), server)
    # 3. 启动server
    server.add_insecure_port('127.0.0.1:50051')
    server.start()
    server.wait_for_termination()
