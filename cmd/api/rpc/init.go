package rpc

import (
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"time"
)

var (
	EtcdAddress    = ""
	MuxConnection  = 1
	Timeout        = 3 * time.Second
	ConnectTimeout = 50 * time.Millisecond
)

func initConfig() error {
	return nil
}

func Init() error {
	if err := initConfig(); err != nil {
		hlog.Fatal("配置初始化失败")
		return err
	}

	if err := initUserRpc(); err != nil {
		return err
	}
	if err := initCommentRpc(); err != nil {
		return err
	}
	if err := initFavoriteRpc(); err != nil {
		return err
	}
	if err := initPublishRpc(); err != nil {
		return err
	}
	if err := initFeedRpc(); err != nil {
		return err
	}
	if err := initMessageRpc(); err != nil {
		return err
	}
	return nil
}
