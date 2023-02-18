package main

import (
	"github.com/cloudwego/kitex/server"
	tracer "github.com/kitex-contrib/tracer-opentracing"
	"log"
	"net"
	"tiktok/config"
	"tiktok/dal/mysql"
	api "tiktok/message/kitex_gen/api/message"
	tracer2 "tiktok/tracer"
)

func main() {
	config.Init()

	tracer2.InitJaeger("message")

	mysql.Init()

	addr, _ := net.ResolveTCPAddr("tcp", config.ServerMsgUrlHP)
	svr := api.NewServer(new(MessageImpl),
		server.WithServiceAddr(addr),
		server.WithSuite(tracer.NewDefaultServerSuite()),
	)

	err := svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
