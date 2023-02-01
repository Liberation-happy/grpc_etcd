package main

import (
	"api-gateway/discovery"
	"api-gateway/internal/service"
	"api-gateway/pkg/util"
	"api-gateway/routes"
	"context"
	"fmt"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
	"google.golang.org/grpc/resolver"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	InitConfig()
	go startListen() //转载路由
	{
		// 初始化一个os.Signal类型channel
		// 我们必须使用缓冲通道，否则在信号发送时如果还没有转杯好接收信号，就有丢失信号的风险
		osSignals := make(chan os.Signal, 1)
		// notify用于监听信号
		// 参数1表示接收信号的channel
		// 参数2及后面的表示要监听的信号
		// os.Interrupt 表示中断
		// os.Kill 杀死退出进程
		signal.Notify(osSignals, os.Interrupt, os.Kill, syscall.SIGTERM, syscall.SIGINT, syscall.SIGKILL)
		// 阻塞直到接收到信息
		s := <-osSignals
		fmt.Println("exit! ", s)
	}
	fmt.Println("gateway listen on :3000")
}

func startListen() {
	// etcd注册
	etcdAddress := []string{viper.GetString("etcd.address")}
	etcdRegister := discovery.NewResolver(etcdAddress, logrus.New())
	resolver.Register(etcdRegister)
	ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)
	// 服务名
	userServiceName := viper.GetString("domain.user")
	taskServiceName := viper.GetString("domain.task")
	pythonServiceName := viper.GetString("domain.python")
	javaServiceName := viper.GetString("domain.java")
	BookServiceName := viper.GetString("domain.book")

	// RPC 连接
	connUser, err := RPCConnect(ctx, userServiceName, etcdRegister)
	if err != nil {
		return
	}
	userService := service.NewUserServiceClient(connUser)

	connTask, err := RPCConnect(ctx, taskServiceName, etcdRegister)
	if err != nil {
		return
	}
	taskService := service.NewTaskServiceClient(connTask)

	connPython, err := RPCConnect(ctx, pythonServiceName, etcdRegister)
	if err != nil {
		return
	}
	PythonService := service.NewGreeterClient(connPython)

	connJava, err := RPCConnect(ctx, javaServiceName, etcdRegister)
	if err != nil {
		return
	}
	JavaService := service.NewJavaHelloServiceClient(connJava)

	connBook, err := RPCConnect(ctx, BookServiceName, etcdRegister)
	if err != nil {
		return
	}
	BookService := service.NewBookControllerClient(connBook)

	// 加入熔断 TODO main太臃肿了
	//wrapper.NewServiceWrapper(userServiceName)
	//wrapper.NewServiceWrapper(taskServiceName)

	ginRouter := routes.NewRouter(userService, taskService, PythonService, JavaService, BookService)
	server := &http.Server{
		Addr:           viper.GetString("server.port"),
		Handler:        ginRouter,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	err = server.ListenAndServe()
	if err != nil {
		fmt.Println("绑定HTTP到 %s 失败！可能是端口已经被占用，或用户权限不足")
		fmt.Println(err)
	}
	go func() {
		// 优雅关闭
		util.GracefullyShutdown(server)
	}()
	if err := server.ListenAndServe(); err != nil {
		fmt.Println("gateway启动失败, err: ", err)
	}
}

func RPCConnect(ctx context.Context, serviceName string, etcdRegister *discovery.Resolver) (conn *grpc.ClientConn, err error) {
	opts := []grpc.DialOption{
		grpc.WithInsecure(),
	}
	addr := fmt.Sprintf("%s:///%s", etcdRegister.Scheme(), serviceName)
	conn, err = grpc.DialContext(ctx, addr, opts...)
	return
}

func InitConfig() {
	workDir, _ := os.Getwd()
	viper.SetConfigName("config")
	viper.SetConfigType("yml")
	viper.AddConfigPath(workDir + "/config")
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
}
