package com.example.grpc;

import static io.grpc.MethodDescriptor.generateFullMethodName;
import static io.grpc.stub.ClientCalls.asyncBidiStreamingCall;
import static io.grpc.stub.ClientCalls.asyncClientStreamingCall;
import static io.grpc.stub.ClientCalls.asyncServerStreamingCall;
import static io.grpc.stub.ClientCalls.asyncUnaryCall;
import static io.grpc.stub.ClientCalls.blockingServerStreamingCall;
import static io.grpc.stub.ClientCalls.blockingUnaryCall;
import static io.grpc.stub.ClientCalls.futureUnaryCall;
import static io.grpc.stub.ServerCalls.asyncBidiStreamingCall;
import static io.grpc.stub.ServerCalls.asyncClientStreamingCall;
import static io.grpc.stub.ServerCalls.asyncServerStreamingCall;
import static io.grpc.stub.ServerCalls.asyncUnaryCall;
import static io.grpc.stub.ServerCalls.asyncUnimplementedStreamingCall;
import static io.grpc.stub.ServerCalls.asyncUnimplementedUnaryCall;

/**
 */
@javax.annotation.Generated(
    value = "by gRPC proto compiler (version 1.28.1)",
    comments = "Source: HiService.proto")
public final class JavaHelloServiceGrpc {

  private JavaHelloServiceGrpc() {}

  public static final String SERVICE_NAME = "com.example.grpc.JavaHelloService";

  // Static method descriptors that strictly reflect the proto.
  private static volatile io.grpc.MethodDescriptor<com.example.grpc.JavaHelloRequest,
      com.example.grpc.JavaHelloResponse> getHelloMethod;

  @io.grpc.stub.annotations.RpcMethod(
      fullMethodName = SERVICE_NAME + '/' + "hello",
      requestType = com.example.grpc.JavaHelloRequest.class,
      responseType = com.example.grpc.JavaHelloResponse.class,
      methodType = io.grpc.MethodDescriptor.MethodType.UNARY)
  public static io.grpc.MethodDescriptor<com.example.grpc.JavaHelloRequest,
      com.example.grpc.JavaHelloResponse> getHelloMethod() {
    io.grpc.MethodDescriptor<com.example.grpc.JavaHelloRequest, com.example.grpc.JavaHelloResponse> getHelloMethod;
    if ((getHelloMethod = JavaHelloServiceGrpc.getHelloMethod) == null) {
      synchronized (JavaHelloServiceGrpc.class) {
        if ((getHelloMethod = JavaHelloServiceGrpc.getHelloMethod) == null) {
          JavaHelloServiceGrpc.getHelloMethod = getHelloMethod =
              io.grpc.MethodDescriptor.<com.example.grpc.JavaHelloRequest, com.example.grpc.JavaHelloResponse>newBuilder()
              .setType(io.grpc.MethodDescriptor.MethodType.UNARY)
              .setFullMethodName(generateFullMethodName(SERVICE_NAME, "hello"))
              .setSampledToLocalTracing(true)
              .setRequestMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.example.grpc.JavaHelloRequest.getDefaultInstance()))
              .setResponseMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.example.grpc.JavaHelloResponse.getDefaultInstance()))
              .setSchemaDescriptor(new JavaHelloServiceMethodDescriptorSupplier("hello"))
              .build();
        }
      }
    }
    return getHelloMethod;
  }

  /**
   * Creates a new async stub that supports all call types for the service
   */
  public static JavaHelloServiceStub newStub(io.grpc.Channel channel) {
    io.grpc.stub.AbstractStub.StubFactory<JavaHelloServiceStub> factory =
      new io.grpc.stub.AbstractStub.StubFactory<JavaHelloServiceStub>() {
        @java.lang.Override
        public JavaHelloServiceStub newStub(io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
          return new JavaHelloServiceStub(channel, callOptions);
        }
      };
    return JavaHelloServiceStub.newStub(factory, channel);
  }

  /**
   * Creates a new blocking-style stub that supports unary and streaming output calls on the service
   */
  public static JavaHelloServiceBlockingStub newBlockingStub(
      io.grpc.Channel channel) {
    io.grpc.stub.AbstractStub.StubFactory<JavaHelloServiceBlockingStub> factory =
      new io.grpc.stub.AbstractStub.StubFactory<JavaHelloServiceBlockingStub>() {
        @java.lang.Override
        public JavaHelloServiceBlockingStub newStub(io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
          return new JavaHelloServiceBlockingStub(channel, callOptions);
        }
      };
    return JavaHelloServiceBlockingStub.newStub(factory, channel);
  }

  /**
   * Creates a new ListenableFuture-style stub that supports unary calls on the service
   */
  public static JavaHelloServiceFutureStub newFutureStub(
      io.grpc.Channel channel) {
    io.grpc.stub.AbstractStub.StubFactory<JavaHelloServiceFutureStub> factory =
      new io.grpc.stub.AbstractStub.StubFactory<JavaHelloServiceFutureStub>() {
        @java.lang.Override
        public JavaHelloServiceFutureStub newStub(io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
          return new JavaHelloServiceFutureStub(channel, callOptions);
        }
      };
    return JavaHelloServiceFutureStub.newStub(factory, channel);
  }

  /**
   */
  public static abstract class JavaHelloServiceImplBase implements io.grpc.BindableService {

    /**
     */
    public void hello(com.example.grpc.JavaHelloRequest request,
        io.grpc.stub.StreamObserver<com.example.grpc.JavaHelloResponse> responseObserver) {
      asyncUnimplementedUnaryCall(getHelloMethod(), responseObserver);
    }

    @java.lang.Override public final io.grpc.ServerServiceDefinition bindService() {
      return io.grpc.ServerServiceDefinition.builder(getServiceDescriptor())
          .addMethod(
            getHelloMethod(),
            asyncUnaryCall(
              new MethodHandlers<
                com.example.grpc.JavaHelloRequest,
                com.example.grpc.JavaHelloResponse>(
                  this, METHODID_HELLO)))
          .build();
    }
  }

  /**
   */
  public static final class JavaHelloServiceStub extends io.grpc.stub.AbstractAsyncStub<JavaHelloServiceStub> {
    private JavaHelloServiceStub(
        io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
      super(channel, callOptions);
    }

    @java.lang.Override
    protected JavaHelloServiceStub build(
        io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
      return new JavaHelloServiceStub(channel, callOptions);
    }

    /**
     */
    public void hello(com.example.grpc.JavaHelloRequest request,
        io.grpc.stub.StreamObserver<com.example.grpc.JavaHelloResponse> responseObserver) {
      asyncUnaryCall(
          getChannel().newCall(getHelloMethod(), getCallOptions()), request, responseObserver);
    }
  }

  /**
   */
  public static final class JavaHelloServiceBlockingStub extends io.grpc.stub.AbstractBlockingStub<JavaHelloServiceBlockingStub> {
    private JavaHelloServiceBlockingStub(
        io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
      super(channel, callOptions);
    }

    @java.lang.Override
    protected JavaHelloServiceBlockingStub build(
        io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
      return new JavaHelloServiceBlockingStub(channel, callOptions);
    }

    /**
     */
    public com.example.grpc.JavaHelloResponse hello(com.example.grpc.JavaHelloRequest request) {
      return blockingUnaryCall(
          getChannel(), getHelloMethod(), getCallOptions(), request);
    }
  }

  /**
   */
  public static final class JavaHelloServiceFutureStub extends io.grpc.stub.AbstractFutureStub<JavaHelloServiceFutureStub> {
    private JavaHelloServiceFutureStub(
        io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
      super(channel, callOptions);
    }

    @java.lang.Override
    protected JavaHelloServiceFutureStub build(
        io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
      return new JavaHelloServiceFutureStub(channel, callOptions);
    }

    /**
     */
    public com.google.common.util.concurrent.ListenableFuture<com.example.grpc.JavaHelloResponse> hello(
        com.example.grpc.JavaHelloRequest request) {
      return futureUnaryCall(
          getChannel().newCall(getHelloMethod(), getCallOptions()), request);
    }
  }

  private static final int METHODID_HELLO = 0;

  private static final class MethodHandlers<Req, Resp> implements
      io.grpc.stub.ServerCalls.UnaryMethod<Req, Resp>,
      io.grpc.stub.ServerCalls.ServerStreamingMethod<Req, Resp>,
      io.grpc.stub.ServerCalls.ClientStreamingMethod<Req, Resp>,
      io.grpc.stub.ServerCalls.BidiStreamingMethod<Req, Resp> {
    private final JavaHelloServiceImplBase serviceImpl;
    private final int methodId;

    MethodHandlers(JavaHelloServiceImplBase serviceImpl, int methodId) {
      this.serviceImpl = serviceImpl;
      this.methodId = methodId;
    }

    @java.lang.Override
    @java.lang.SuppressWarnings("unchecked")
    public void invoke(Req request, io.grpc.stub.StreamObserver<Resp> responseObserver) {
      switch (methodId) {
        case METHODID_HELLO:
          serviceImpl.hello((com.example.grpc.JavaHelloRequest) request,
              (io.grpc.stub.StreamObserver<com.example.grpc.JavaHelloResponse>) responseObserver);
          break;
        default:
          throw new AssertionError();
      }
    }

    @java.lang.Override
    @java.lang.SuppressWarnings("unchecked")
    public io.grpc.stub.StreamObserver<Req> invoke(
        io.grpc.stub.StreamObserver<Resp> responseObserver) {
      switch (methodId) {
        default:
          throw new AssertionError();
      }
    }
  }

  private static abstract class JavaHelloServiceBaseDescriptorSupplier
      implements io.grpc.protobuf.ProtoFileDescriptorSupplier, io.grpc.protobuf.ProtoServiceDescriptorSupplier {
    JavaHelloServiceBaseDescriptorSupplier() {}

    @java.lang.Override
    public com.google.protobuf.Descriptors.FileDescriptor getFileDescriptor() {
      return com.example.grpc.HiService.getDescriptor();
    }

    @java.lang.Override
    public com.google.protobuf.Descriptors.ServiceDescriptor getServiceDescriptor() {
      return getFileDescriptor().findServiceByName("JavaHelloService");
    }
  }

  private static final class JavaHelloServiceFileDescriptorSupplier
      extends JavaHelloServiceBaseDescriptorSupplier {
    JavaHelloServiceFileDescriptorSupplier() {}
  }

  private static final class JavaHelloServiceMethodDescriptorSupplier
      extends JavaHelloServiceBaseDescriptorSupplier
      implements io.grpc.protobuf.ProtoMethodDescriptorSupplier {
    private final String methodName;

    JavaHelloServiceMethodDescriptorSupplier(String methodName) {
      this.methodName = methodName;
    }

    @java.lang.Override
    public com.google.protobuf.Descriptors.MethodDescriptor getMethodDescriptor() {
      return getServiceDescriptor().findMethodByName(methodName);
    }
  }

  private static volatile io.grpc.ServiceDescriptor serviceDescriptor;

  public static io.grpc.ServiceDescriptor getServiceDescriptor() {
    io.grpc.ServiceDescriptor result = serviceDescriptor;
    if (result == null) {
      synchronized (JavaHelloServiceGrpc.class) {
        result = serviceDescriptor;
        if (result == null) {
          serviceDescriptor = result = io.grpc.ServiceDescriptor.newBuilder(SERVICE_NAME)
              .setSchemaDescriptor(new JavaHelloServiceFileDescriptorSupplier())
              .addMethod(getHelloMethod())
              .build();
        }
      }
    }
    return result;
  }
}
