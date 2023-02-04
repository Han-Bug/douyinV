package main

import (
	"github.com/cloudwego/kitex/server"
	"log"
	"net"
	"tiktok/config"
	api "tiktok/feed/kitex_gen/api/feed"
)

func main() {
	config.Init()

	addr, _ := net.ResolveTCPAddr("tcp", config.ServerFeedUrlHP)
	svr := api.NewServer(new(FeedImpl), server.WithServiceAddr(addr))

	err := svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
