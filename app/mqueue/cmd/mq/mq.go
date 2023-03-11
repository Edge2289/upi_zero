package main

import (
	"flag"
	"upi/app/mqueue/cmd/mq/internal/config"
	"upi/app/mqueue/cmd/mq/internal/listen"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/service"
)

var configFile = flag.String("f", "etc/mq.yaml", "Specify the config file")

func main()  {
	flag.Parse()
	var c config.Config

	conf.MustLoad(*configFile, &c)

	if err := c.SetUp(); err != nil {
		panic(err)
	}

	serviceGroup := service.NewServiceGroup()
	defer serviceGroup.Stop()

	for _, mq := range listen.Mqs(c){
		serviceGroup.Add(mq)
	}

	serviceGroup.Start()
}