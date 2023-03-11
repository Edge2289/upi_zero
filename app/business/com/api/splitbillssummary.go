package main

import (
	"flag"
	"fmt"

	"upi/app/business/com/api/internal/config"
	"upi/app/business/com/api/internal/handler"
	"upi/app/business/com/api/internal/svc"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/rest"

	"upi/common/middleware"
)

var configFile = flag.String("f", "etc/splitbillssummary.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)

	server := rest.MustNewServer(c.RestConf, middleware.BaseMiddleware()...)
	defer server.Stop()

	ctx := svc.NewServiceContext(c)
	handler.RegisterHandlers(server, ctx)

	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	server.Start()
}
