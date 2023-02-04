package main

import (
	"github.com/cloudwego/kitex/server"
	"log"
	"net"
	"tiktok/config"
	"tiktok/dal/mysql"
	api "tiktok/message/kitex_gen/api/message"
)

func main() {
	config.Init()

	mysql.Init()

	addr, _ := net.ResolveTCPAddr("tcp", config.ServerMsgUrlHP)
	svr := api.NewServer(new(MessageImpl), server.WithServiceAddr(addr))

	err := svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
