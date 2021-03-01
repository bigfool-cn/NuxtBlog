package middlewares

import (
	"github.com/gin-gonic/gin"
	"nuxt-blog-api/models"
)

func VistorMiddleWare() gin.HandlerFunc  {
	return func(c *gin.Context) {
		var vistorModel = models.Vistor{
			VistorIP: c.ClientIP(),
			UserAgent: c.Request.UserAgent(),
			VistorPath: c.Request.URL.Path,
		}
		go vistorModel.Create()
	}
}