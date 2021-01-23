package routers

import (
	"github.com/gin-gonic/gin"
	"nuxt-blog-api/apis"
	"nuxt-blog-api/middlewares"
)

func InitRouter() *gin.Engine  {
	r := gin.New()

	// r.Use(Cors())
	r.Use(middlewares.VistorMiddleWare())

	r.POST("/login", apis.Login)

	auth := r.Group("/admin")
	auth.Use(middlewares.JwtMiddleWare())
	{
		// 访客
		vistor := auth.Group("vistors")
		{
			vistor.GET("", apis.VistorList)
			vistor.DELETE("", apis.DeleteVistor)
		}

		// 标签管理
		tag := auth.Group("tags")
		{
			tag.POST("", apis.CreateTag)
			tag.PUT("/:tag_id", apis.UpdateTag)
			tag.GET("", apis.GetTags)
			tag.DELETE("", apis.DeleteTag)
		}

		// 文章管理
		article := auth.Group("articles")
		{
			article.GET("", apis.ArticleList)
			article.GET("/:article_id", apis.GetArticle)
			article.POST("", apis.CreateArticle)
			article.PUT("/:article_id", apis.UpdateArticle)
			article.DELETE("", apis.DeleteArticle)
		}

	}

	r.GET("/tags", apis.GetBlogTags)
	r.GET("/articles/:articleId", apis.GetBlogArticle)
	r.GET("/articles", apis.GetBlogArticleList)
	return r
}


func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		method := c.Request.Method
		c.Header("Access-Control-Allow-Origin", "*")  // 可将将 * 替换为指定的域名
		c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE, UPDATE")
		c.Header("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept, Authorization")
		c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Cache-Control, Content-Language, Content-Type")
		c.Header("Access-Control-Allow-Credentials", "true")
		if method == "OPTIONS" {
			c.AbortWithStatus(204)
		}
		c.Next()
	}
}
