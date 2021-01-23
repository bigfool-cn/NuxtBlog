package middlewares

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"nuxt-blog-api/apis"
	"nuxt-blog-api/utils"
	"time"
)

// 验证jwt令牌
func JwtMiddleWare() gin.HandlerFunc  {
	return func(c *gin.Context) {
		token := c.GetHeader("Authorization")
		if token == "" {
			c.JSON(http.StatusOK,apis.Response{Code:401,Message:"token不可用"})
			c.Abort()
			return
		} else {
			claims, err := utils.Jwt.ParseToken(token)
			if err != nil {
				log.Println(err.Error())
				c.JSON(http.StatusOK,apis.Response{Code:401,Message:"token不可用"})
				c.Abort()
				return
			} else if time.Now().Unix() > claims.ExpiresAt {
				c.JSON(http.StatusOK,apis.Response{Code:401,Message:"token已过期"})
				c.Abort()
				return
			}
			c.Set("_user_name",claims.UserName)
		}
		c.Next()
	}
}
