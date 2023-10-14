package _jwt

import (
	"errors"
	"github.com/golang-jwt/jwt/v5"
	"ocean_mall/account/conf"
	account_log "ocean_mall/account/log"
	"time"
)

var (
	TokenExpired     = errors.New("token 已过期")
	TokenNotValidYet = errors.New("token 不在有效")
	TokenMalformed   = errors.New("token 非法")
	TokenInvalid     = errors.New("token 不可用")
)

type CustomClaims struct {
	AccountId   int32
	NickName    string
	AuthorityId int32
	jwt.RegisteredClaims
}

type JWT struct {
	SignKey []byte
}

func NewJWT() *JWT {
	return &JWT{
		SignKey: []byte(conf.AppConf.JWTConfig.SignKey),
	}
}

func (j *JWT) GenerateJWT(claims CustomClaims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenStr, err := token.SignedString([]byte(conf.AppConf.JWTConfig.SignKey))
	if err != nil {
		panic(err)
		account_log.Logger.Error("生成 JWT 错误" + err.Error())
		return "", err
	}
	return tokenStr, nil

}

func (j *JWT) ParseJWT(tokenStr string) (*CustomClaims, error) {
	token, _ := jwt.ParseWithClaims(tokenStr, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return j.SignKey, nil
	})

	if token != nil {
		if claims, ok := token.Claims.(*CustomClaims); ok && token.Valid {
			return claims, nil
		}

	}
	return nil, TokenInvalid
}

func (j *JWT) RefreshJWT(tokenStr string) (string, error) {
	jwt.WithTimeFunc(func() time.Time {
		return time.Unix(0, 0)
	})

	token, err := jwt.ParseWithClaims(tokenStr, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return j.SignKey, nil
	})

	if err != nil {
		panic(err)
	}

	if claims, ok := token.Claims.(*CustomClaims); ok && token.Valid {
		jwt.WithTimeFunc(func() time.Time {
			return time.Now()
		})
		claims.RegisteredClaims.ExpiresAt = jwt.NewNumericDate(time.Now().Add(7 * 24 * time.Hour))
		return j.GenerateJWT(*claims)
	}

	return "", TokenInvalid

}
