package global

import (
	"douyinV/api/config"
	"douyinV/api/kitex_gen/user/usersvr"
	"douyinV/api/pkg/jwt"
)

var APIImpl ApiImpl

type ApiImpl struct {
	ApiConfig  config.APIConfig
	JWTImpl    jwt.JwtImpl
	UserClient usersvr.Client
}
