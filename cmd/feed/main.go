package main

import (
	"github.com/cloudwego/kitex/server"
	tracer "github.com/kitex-contrib/tracer-opentracing"
	"log"
	"net"
	"tiktok/config"
	api "tiktok/feed/kitex_gen/api/feed"
	tracer2 "tiktok/tracer"
)

func main() {
	config.Init()

	tracer2.InitJaeger("feed")

	addr, _ := net.ResolveTCPAddr("tcp", config.ServerFeedUrlHP)
	svr := api.NewServer(new(FeedImpl),
		server.WithServiceAddr(addr),
		server.WithSuite(tracer.NewDefaultServerSuite()),
	)

	err := svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
