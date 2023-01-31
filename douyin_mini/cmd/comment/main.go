package main

import (
	main2 "douyin_mini/cmd/comment"
	comment "douyin_mini/kitex_gen/comment/commentservice"
	"log"
)

func main() {
	svr := comment.NewServer(new(main2.CommentServiceImpl))

	err := svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
