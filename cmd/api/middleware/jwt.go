package middleware

import "github.com/hertz-contrib/jwt"

var (
	JwtMiddleware *jwt.HertzJWTMiddleware
	IdentityKey   = "identity"
)

func InitJWT() {
	//var err error

}
