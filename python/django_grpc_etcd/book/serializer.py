from .utils import generics, proto_serializers
from .proto_py import book_pb2, book_pb2_grpc
from .models import Book


class BookProtoSerializer(proto_serializers.ModelProtoSerializer):
    class Meta:
        model = Book
        proto_class = book_pb2.Book
        fields = ["id", "title", "comment"]
