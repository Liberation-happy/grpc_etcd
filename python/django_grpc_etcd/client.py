import grpc
from book.proto_py import book_pb2, book_pb2_grpc

with grpc.insecure_channel('localhost:10003') as channel:
    stub = book_pb2_grpc.BookControllerStub(channel)
    for book in stub.List(book_pb2.BookListRequest()):
        print(book, end='')
