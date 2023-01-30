import com.example.grpc.JavaHelloRequest;
import com.example.grpc.JavaHelloResponse;
import com.example.grpc.JavaHelloServiceGrpc;
import io.grpc.Server;
import io.grpc.ServerBuilder;
import io.grpc.stub.StreamObserver;

import java.io.IOException;

public class MyGrpcServer extends JavaHelloServiceGrpc.JavaHelloServiceImplBase {
    static public void main(String[] args) throws InterruptedException, IOException {
        Server server = ServerBuilder.forPort(8082)
                .addService(new MyGrpcServer())
                .build();
        System.out.println("Starting server");
        Register register = new Register("http://localhost:2379");
        String key = "/java_server/127.0.0.1:8082";
        String value = "{\"name\":\"java_server\",\"addr\":\"127.0.0.1:8082\",\"version\":\"\",\"weight\": 0}";
        try {
            register.putWithLease(key, value);
        } catch (Exception e) {
            e.printStackTrace();
        }
        server.start();
        System.out.println("Server started");
        server.awaitTermination();
    }

    public void hello(JavaHelloRequest request, StreamObserver<JavaHelloResponse> responseObserver) {
        System.out.println(1);
        System.out.println(request);

        String greeting = "Hi " + request.getName() + " you are " + request.getAge() + " years old " +
                "your hobby is " + (request.getHobbiesList()) + " your tags " + request.getTagsMap();

        JavaHelloResponse response = JavaHelloResponse.newBuilder().setGreeting(greeting).build();
        responseObserver.onNext(response);
        responseObserver.onCompleted();
    }
}
