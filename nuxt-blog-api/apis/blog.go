package apis

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"nuxt-blog-api/cache"
	"nuxt-blog-api/models"
	"strconv"
)

// 获取所有标签
func GetBlogTags (c *gin.Context){
	var tagModel models.Tag

	tagsStr := cache.RedisCache.Get("blog_tags")
	if tagsStr != "" {
		var tags []models.Tag
		err := json.Unmarshal([]byte(tagsStr),&tags)
		if err == nil {
			c.JSON(http.StatusOK,Response{Code:200,Message:"获取成功",Data:tags})
			return
		}
	}

	tags, err := tagModel.GetTags()

	if err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusOK,Response{Code:400,Message:"获取数据失败"})
		return
	}

	tagsByte, err := json.Marshal(tags)
	if err != nil {
		c.JSON(http.StatusOK,Response{Code:400,Message:"获取数据失败"})
		return
	}
	cache.RedisCache.Set("blog_tags" , string(tagsByte),0)

	c.JSON(http.StatusOK,Response{Code:200,Message:"获取成功",Data:tags})
}

// 文章列表
func GetBlogArticleList(c *gin.Context)  {
	var articleModel models.Article

	pageIndex, _ := strconv.Atoi(c.DefaultQuery("page","1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("limit","20"))
	tagId, _ := strconv.ParseInt(c.DefaultQuery("tag_id","-1"),10,64)

	_articles, total, err := articleModel.GetArticlePage(pageSize,pageIndex,"",1,tagId)
	if err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusOK,Response{Code:400,Message:"获取数据失败"})
		return
	}
	type articles struct {
		Articles  []models.Article `json:"articles"`
		Total     int64            `json:"total"`
	}
	c.JSON(http.StatusOK,Response{Code:200,Message:"获取成功",Data:&articles{Articles:_articles,Total:total}})
}

// 获取文章
func GetBlogArticle(c *gin.Context)  {
	var articleModel models.Article
	
	articleId, _ := strconv.ParseInt(c.Param("articleId"),10,64)

	if articleId == 0{
		c.JSON(http.StatusOK,Response{Code:400,Message:"参数错误"})
		return
	}
	articleModel.ArticleID = articleId
	article, err := articleModel.GetArticle()
	if err != nil {
		c.JSON(http.StatusOK,Response{Code:400,Message:"获取数据失败"})
		return
	}
	article.ArticleRead = article.ArticleRead + 1
	if err := articleModel.UpdateArticleRead(articleId,article.ArticleRead); err != nil {
		c.JSON(http.StatusOK,Response{Code:400,Message:"获取数据失败"})
		return
	}
	c.JSON(http.StatusOK,Response{Code:200,Message:"获取成功",Data:article})
}
