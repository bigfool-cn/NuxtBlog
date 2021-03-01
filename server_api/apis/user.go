package apis

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"nuxt-blog-api/cache"
	"nuxt-blog-api/configs"
	"nuxt-blog-api/utils"
	"strconv"
	"time"
)

type user struct {
	Username  string `json:"username" binding:"required"`
	Password  string `json:"password"" binding:"required"`
 }

func Login(c *gin.Context)  {
	var user user
	err := c.BindJSON(&user)

	// 登录限制
	limit, _ := strconv.ParseInt(cache.RedisCache.Get(c.ClientIP()),10,64)
	if limit == 0 {
		cache.RedisCache.Set(c.ClientIP(),"0", 24 * time.Hour)
	}
	if limit >= 5 {
		c.JSON(http.StatusOK,Response{Code: 400,Message: "帐号已被限制登录"})
		return
	}

	if err != nil {
		cache.RedisClient.Incr(c.ClientIP())
		c.JSON(http.StatusOK,Response{Code: 400,Message: err.Error()})
		return
	}

	if user.Username != configs.Configs.AdminUser.UserName {
		cache.RedisClient.Incr(c.ClientIP())
		c.JSON(http.StatusOK,Response{Code: 400,Message: "用户名不正确"})
		return
	}

	pwd := utils.MD5(user.Password)
	if pwd != configs.Configs.AdminUser.Pwd {
		cache.RedisClient.Incr(c.ClientIP())
		c.JSON(http.StatusOK,Response{Code: 400,Message: "密码不正确"})
		return
	}

	token, err := utils.Jwt.GenerateToken(configs.Configs.AdminUser.UserName)
	if err != nil {
		c.JSON(http.StatusOK,Response{Code: 400,Message: "获取token失败"})
		return
	}

	c.JSON(http.StatusOK,Response{Code: 200,Message: "登录成功",Data: token})
	return
}
