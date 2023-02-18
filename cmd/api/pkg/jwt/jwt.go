package jwt

import (
	"errors"
	"github.com/golang-jwt/jwt/v4"
	"time"
)

// JwtImpl 对JWT的简单自定义封装
type JwtImpl struct {
	SecretKey []byte
	ValidTime time.Duration
	Issuer    string
	Subject   string
}

// myClaim Data为token存储的数据
type myClaim struct {
	Data TokenData
	jwt.RegisteredClaims
}
type TokenData struct {
	UserId int64
}

func (j *JwtImpl) MakeToken(tokenData TokenData) (tokenString string, err error) {
	claim := myClaim{
		Data: tokenData,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:  j.Issuer,
			Subject: j.Subject,
			// 过期时间
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(j.ValidTime)),
			// 生效时间
			NotBefore: jwt.NewNumericDate(time.Now()),
			// 签发时间
			IssuedAt: jwt.NewNumericDate(time.Now()),
			ID:       "",
		},
	}
	// 使用HS256算法，声明token对象  **使用ES256时密文用byte[]是无效的
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)
	// 生成token字符串（根据密文生成）
	tokenString, err = token.SignedString(j.SecretKey)
	return tokenString, err
}

func (j *JwtImpl) ParseToken(tokenStr string) (TokenData, error) {
	token, err := jwt.ParseWithClaims(tokenStr, &myClaim{}, func(token *jwt.Token) (interface{}, error) {
		return j.SecretKey, nil
	})

	if err != nil {
		// 如果是jwt内部传出的校验错误
		if ve, ok := err.(*jwt.ValidationError); ok {
			if ve.Errors&jwt.ValidationErrorMalformed != 0 {
				return TokenData{}, errors.New("不是符合格式的token字符串")
			} else if ve.Errors&jwt.ValidationErrorExpired != 0 {
				return TokenData{}, errors.New("过期的token")
			} else if ve.Errors&jwt.ValidationErrorNotValidYet != 0 {
				return TokenData{}, errors.New("无效的token")
			} else {
				return TokenData{}, errors.New("无法处理该token（jwt内部意外的错误）")
			}
		}
	}
	// 如果能从token中解析成自定义token数据，且token是有效的
	if claims, ok := token.Claims.(*myClaim); ok && token.Valid {
		return claims.Data, nil
	}
	return TokenData{}, errors.New("无法处理该token（jwt外部意外的错误）")
}
