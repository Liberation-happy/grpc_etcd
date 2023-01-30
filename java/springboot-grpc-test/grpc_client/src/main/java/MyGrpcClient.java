import com.example.grpc.JavaHelloRequest;
import com.example.grpc.JavaHelloResponse;
import com.example.grpc.JavaHelloServiceGrpc;
import io.grpc.ManagedChannel;
import io.grpc.ManagedChannelBuilder;

public class MyGrpcClient {
    public static void main(String[] args) throws InterruptedException {
        ManagedChannel channel = ManagedChannelBuilder.forAddress("localhost", 8082)
                .usePlaintext()
                .build();

        JavaHelloServiceGrpc.JavaHelloServiceBlockingStub stub =
                JavaHelloServiceGrpc.newBlockingStub(channel);

        JavaHelloResponse helloResponse = stub.hello(
                JavaHelloRequest.newBuilder()
                        .setName("xxz")
                        .setAge(20)
                        .addHobbies("basketball").putTags("how?", "wonderful")
                        .build());

        System.out.println(helloResponse);

        channel.shutdown();

    }
}
