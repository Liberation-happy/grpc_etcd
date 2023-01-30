from .models import Book
from .proto_py import helloworld_pb2, helloworld_pb2_grpc, book_pb2_grpc, book_pb2
from .serializer import BookProtoSerializer
from .utils import generics


class HelloWorldServer(helloworld_pb2_grpc.GreeterServicer):
    def SayHello(self, request, context):
        print(request.name)
        return helloworld_pb2.HelloResponse(message="nihao")

    def start_server(self, server):
        helloworld_pb2_grpc.add_GreeterServicer_to_server(HelloWorldServer(), server)


class BookService(generics.ModelService):
    queryset = Book.objects.all()
    serializer_class = BookProtoSerializer

    def book_server(self, server):
        book_pb2_grpc.add_BookControllerServicer_to_server(BookService.as_servicer(), server)
