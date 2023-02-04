package config

import (
	"github.com/joho/godotenv"
	"os"
	"path"
	"runtime"
)

var (
	MysqlDsn        string
	ClientUrlHP     string
	ClientUrl       string
	ServerFeedUrlHP string
	ServerMsgUrlHP  string
)

func Init() {
	if err := godotenv.Load(getCurrentAbPath() + "/.env"); err != nil {
		panic("读取配置失败！" + err.Error())
	}

	MysqlDsn = os.Getenv("MYSQL_DSN")
	ClientUrlHP = os.Getenv("CLIENT_URL_HP")
	ClientUrl = os.Getenv("CLIENT_URL_PROTOCOL") + "://" + ClientUrlHP
	ServerFeedUrlHP = os.Getenv("SERVER_FEED_URL_HP")
	ServerMsgUrlHP = os.Getenv("SERVER_MSG_URL_HP")
}

func getCurrentAbPath() string {
	var abPath string
	_, filename, _, ok := runtime.Caller(0)
	if ok {
		abPath = path.Dir(filename)
	}
	return abPath
}
