import io.etcd.jetcd.ByteSequence;
import io.etcd.jetcd.Client;
import io.etcd.jetcd.KV;
import io.etcd.jetcd.Lease;
import io.etcd.jetcd.Response;
import io.etcd.jetcd.lease.LeaseKeepAliveResponse;
import io.etcd.jetcd.options.PutOption;
import io.grpc.stub.CallStreamObserver;

import static com.google.common.base.Charsets.UTF_8;

public class Register {
    private volatile Client client;
    private final String endpoints;
    private final Object lock = new Object();


    public Register(String endpoints) {
        super();
        this.endpoints = endpoints;
    }

    /**
     * 新建key-value客户端实例
     * @return
     */
    private KV getKVClient(){

        if (null==client) {
            synchronized (lock) {
                if (null==client) {

                    client = Client.builder().endpoints(endpoints.split(",")).build();
                }
            }
        }

        return client.getKVClient();
    }

    public void close() {
        client.close();
        client = null;
    }

    public Response.Header put(String key, String value) throws Exception {
        return getKVClient().put(bytesOf(key), bytesOf(value)).get().getHeader();
    }

    /**
     * 将字符串转为客户端所需的ByteSequence实例
     * @param val
     * @return
     */
    public static ByteSequence bytesOf(String val) {
        return ByteSequence.from(val, UTF_8);
    }

    private Client getClient() {
        if (null==client) {
            synchronized (lock) {
                if (null==client) {
                    client = Client.builder().endpoints(endpoints.split(",")).build();
                }
            }
        }

        return client;
    }

    public void putWithLease(String key, String value) throws Exception {
        Lease leaseClient = getClient().getLeaseClient();

        leaseClient.grant(60).thenAccept(result -> {
            // 租约ID
            long leaseId = result.getID();

            // 准备好put操作的client
            KV kvClient = getClient().getKVClient();

            // put操作时的可选项，在这里指定租约ID
            PutOption putOption = PutOption.newBuilder().withLeaseId(leaseId).build();

            // put操作
            kvClient.put(bytesOf(key), bytesOf(value), putOption)
                    .thenAccept(putResponse -> {
                        // put操作完成后，再设置无限续租的操作
                        leaseClient.keepAlive(leaseId, new CallStreamObserver<LeaseKeepAliveResponse>() {
                            @Override
                            public boolean isReady() {
                                return false;
                            }

                            @Override
                            public void setOnReadyHandler(Runnable onReadyHandler) {

                            }

                            @Override
                            public void disableAutoInboundFlowControl() {

                            }

                            @Override
                            public void request(int count) {
                            }

                            @Override
                            public void setMessageCompression(boolean enable) {

                            }

                            /**
                             * 每次续租操作完成后，该方法都会被调用
                             * @param value
                             */
                            @Override
                            public void onNext(LeaseKeepAliveResponse value) {
                                System.out.println("续租完成");
                            }

                            @Override
                            public void onError(Throwable t) {
                                System.out.println(t);
                            }

                            @Override
                            public void onCompleted() {
                            }
                        });
                    });
        });
    }
}
