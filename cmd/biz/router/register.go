// Code generated by hertz generator. DO NOT EDIT.

package router

import (
	"github.com/cloudwego/hertz/pkg/app/server"
	feed "tiktok/biz/router/feed"
	message "tiktok/biz/router/message"
)

// GeneratedRegister registers routers generated by IDL.
func GeneratedRegister(r *server.Hertz) {
	//INSERT_POINT: DO NOT DELETE THIS LINE!
	feed.Register(r)

	message.Register(r)
}