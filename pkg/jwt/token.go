package jwt

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"time"
)

var accessTokenSecret = []byte("accessTokenSecret")
var refreshTokenSecret = []byte("refreshTokenSecret")

type Claims struct {
	UserName string
	PassWord string
	jwt.StandardClaims
}

// GenerateToken 生成令牌
//accessTokenString(AT)，refreshTokenString(RT).
//当AT失效且RT有效时才会重新生产AT,RT.当AT和RT都失效时才会让用户重新登录.当AT未失效不会使用到RT.
func GenerateToken(userName, passWord string) (string, string, error) {
	shortExpireTime := time.Now().Add(30 * time.Minute)
	longExpireTime := time.Now().Add(7 * 24 * time.Hour)
	accessTokenClaims := &Claims{
		UserName: userName,
		PassWord: passWord,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: shortExpireTime.Unix(),
			IssuedAt:  time.Now().Unix(),
			Issuer:    "127.0.0.1",
			Subject:   "user token",
		},
	}
	refreshTokenClaims := &Claims{
		UserName: userName,
		PassWord: passWord,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: longExpireTime.Unix(),
			IssuedAt:  time.Now().Unix(),
			Issuer:    "127.0.0.1",
			Subject:   "user token",
		},
	}

	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, accessTokenClaims)
	accessTokenString, err := accessToken.SignedString(accessTokenSecret)
	if err != nil {
		fmt.Println(err)
		return "", "", err
	}

	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshTokenClaims)
	refreshTokenString, err := refreshToken.SignedString(refreshTokenSecret)
	if err != nil {
		fmt.Println(err)
		return "", "", err
	}

	return accessTokenString, refreshTokenString, nil
}

// ParseToken 解析Token
func ParseToken(ctx *gin.Context) {

}
