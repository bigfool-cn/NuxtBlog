package utils

import (
	jwt "github.com/dgrijalva/jwt-go"
	"nuxt-blog-api/configs"
	"strings"
	"time"
)

var jwtSecret = []byte(configs.Configs.App.JwtSecret)

type Claims struct {
	UserName string  `json:"username"`
	jwt.StandardClaims
}

type JwtDo struct {
}

// 生成token
func (JwtDo) GenerateToken(userName string) (string, error) {
	nowTime := time.Now()
	expireTime := nowTime.Add(8 * time.Hour)
	claims := Claims{
		userName,
		jwt.StandardClaims {
			ExpiresAt : expireTime.Unix(),
			Issuer : "bigfool.cn",
		},
	}

	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := tokenClaims.SignedString(jwtSecret)

	return token, err
}

// 解析token
func (JwtDo) ParseToken(token string) (*Claims, error) {
	if strings.Contains(token,"Bearer ") {
		token = strings.Replace(token,"Bearer ","",1)
	}
	tokenClaims, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})

	if tokenClaims != nil {
		if claims, ok := tokenClaims.Claims.(*Claims); ok && tokenClaims.Valid {
			return claims, nil
		}
	}

	return nil, err
}

var Jwt = new(JwtDo)
