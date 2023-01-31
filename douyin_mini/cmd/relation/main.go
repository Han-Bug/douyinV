package main

import (
	"douyin_mini/cmd/relation/handlers"
	relation "douyin_mini/kitex_gen/relation/relationservice"
	"log"
)

func main() {
	svr := relation.NewServer(new(handlers.RelationServiceImpl))

	err := svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
