package main

import (
	"flag"
	"fmt"
	"upi/common/interceptor/rpcserver"

	"upi/app/unionpay/com/rpc/internal/config"
	"upi/app/unionpay/com/rpc/internal/server"
	"upi/app/unionpay/com/rpc/internal/svc"
	"upi/app/unionpay/com/rpc/pb"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/service"
	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var configFile = flag.String("f", "etc/unionpay.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)
	ctx := svc.NewServiceContext(c)

	s := zrpc.MustNewServer(c.RpcServerConf, func(grpcServer *grpc.Server) {
		pb.RegisterUnionpayServer(grpcServer, server.NewUnionpayServer(ctx))
		pb.RegisterStreamGreeterServer(grpcServer, server.NewStreamGreeterServer(ctx))

		if c.Mode == service.DevMode || c.Mode == service.TestMode {
			reflection.Register(grpcServer)
		}
	})
	//rpc log
	s.AddUnaryInterceptors(rpcserver.LoggerInterceptor)
	defer s.Stop()

	fmt.Printf("Starting rpc server at %s...\n", c.ListenOn)
	s.Start()
}
