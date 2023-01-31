package main

import (
	main2 "douyin_mini/cmd/video/handlers"
	video "douyin_mini/kitex_gen/video/videoservice"
	"log"
)

func main() {
	svr := video.NewServer(new(main2.VideoServiceImpl))

	err := svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
