package main

import (
	"CloudDIsk/dataService4/mq"
	"flag"
	"fmt"

	"CloudDIsk/dataService4/internal/config"
	"CloudDIsk/dataService4/internal/handler"
	"CloudDIsk/dataService4/internal/svc"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/rest"
)

var configFile = flag.String("f", "etc/dataservice-api.yaml", "the config file")

func main() {
	flag.Parse()

	//var c config.Config
	conf.MustLoad(*configFile, &config.Conf)

	server := rest.MustNewServer(config.Conf.RestConf)
	defer server.Stop()

	ctx := svc.NewServiceContext(config.Conf)
	handler.RegisterHandlers(server, ctx)

	fmt.Printf("Starting server at %s:%d...\n", config.Conf.Host, config.Conf.Port)
	go mq.StartHeartbeat()
	go mq.StartLocate()
	server.Start()
}