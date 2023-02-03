package config

import (
	"github.com/joho/godotenv"
	"os"
	"path"
	"runtime"
)

var (
	MysqlDsn              string
	ClientUrlHostNamePORT string
	ClientUrl             string
	ServerUrlHostNamePORT string
	ServerUrl             string
)

func Init() {
	if err := godotenv.Load(getCurrentAbPath() + "/.env"); err != nil {
		panic("读取配置失败！" + err.Error())
	}

	MysqlDsn = os.Getenv("MYSQL_DSN")
	ClientUrlHostNamePORT = os.Getenv("CLIENT_URL_HOSTNAME_PORT")
	ClientUrl = os.Getenv("CLIENT_URL_PROTOCOL") + "://" + ClientUrlHostNamePORT
	ServerUrlHostNamePORT = os.Getenv("SERVER_URL_HOSTNAME_PORT")
	ServerUrl = os.Getenv("SERVER_URL_PROTOCOL") + "://" + ClientUrlHostNamePORT
}

func getCurrentAbPath() string {
	var abPath string
	_, filename, _, ok := runtime.Caller(0)
	if ok {
		abPath = path.Dir(filename)
	}
	return abPath
}
