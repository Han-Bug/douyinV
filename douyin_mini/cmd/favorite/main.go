package main

import (
	favorite2 "douyin_mini/cmd/favorite/handlers"
	favorite "douyin_mini/kitex_gen/favorite/favoriteservice"
	"log"
)

func main() {
	svr := favorite.NewServer(new(favorite2.FavoriteServiceImpl))

	err := svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
