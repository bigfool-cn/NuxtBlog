package apis

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"nuxt-blog-api/models"
	"strconv"
)

// 获取文章
func GetArticle(c *gin.Context)  {
	var articleModel models.Article

	articleId, err := strconv.ParseInt(c.Param("article_id"),10,64)
	if err != nil {
		c.JSON(http.StatusOK,Response{Code:400,Message:"参数验证失败"})
		return
	}

	articleModel.ArticleID = articleId

	article, err := articleModel.GetArticle()
	if err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusOK,Response{Code:400,Message:"获取失败"})
		return
	}

	c.JSON(http.StatusOK,Response{Code:200,Message:"获取成功",Data:article})
}

//添加文章
func CreateArticle(c *gin.Context)  {
	var (
		articleForm models.ArticleForm
		articleModel models.Article
	)
	if err := c.BindJSON(&articleForm); err != nil {
		log.Println(err.Error())
		c.JSON(400,Response{Code:400,Message:"数据验证失败," + err.Error()})
		return
	}

	articleModel.ArticleTitle = articleForm.ArticleTitle
	articleModel.ArticleStatus = articleForm.ArticleStatus
	articleModel.ArticleDesc = articleForm.ArticleDesc
	articleModel.ArticleContent = articleForm.ArticleContent

	for _, tagId := range articleForm.TagIds {
		tag := models.Tag{
			TagID:     tagId,
		}
		articleModel.Tags = append(articleModel.Tags,tag)
	}

	if err := articleModel.Create(); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusOK,Response{Code:400,Message:"添加失败"})
		return
	}
	c.JSON(http.StatusOK,Response{Code:200,Message:"添加成功"})
}

// 修改文章
func UpdateArticle(c *gin.Context)  {
	var (
		articleForm models.ArticleForm
		articleModel models.Article
	)
	if err := c.BindJSON(&articleForm); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusOK,Response{Code:400,Message:"数据验证失败," + err.Error()})
		return
	}

	articleId, err := strconv.ParseInt(c.Param("article_id"),10,64)
	if  err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusOK,Response{Code:400,Message:"参数验证失败"})
		return
	}

	if article, err := articleModel.GetArticleById(articleId); article.ArticleID == 0 || err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusOK,Response{Code:400,Message:"文章不存在"})
		return
	}

	articleModel.ArticleID = articleId
	articleModel.ArticleTitle = articleForm.ArticleTitle
	articleModel.ArticleStatus = articleForm.ArticleStatus
	articleModel.ArticleDesc = articleForm.ArticleDesc
	articleModel.ArticleContent = articleForm.ArticleContent

	for _, tagId := range articleForm.TagIds {
		tag := models.Tag{
			TagID:     tagId,
		}
		articleModel.Tags = append(articleModel.Tags,tag)
	}

	if err := articleModel.Update(); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusOK,Response{Code:400,Message:"修改失败"})
		return
	}
	c.JSON(http.StatusOK,Response{Code:200,Message:"修改成功"})
}

// 删除文章
func DeleteArticle(c *gin.Context)  {
	var (
		articleIds   []int64
		articleModel models.Article
	)

	if err := c.BindJSON(&articleIds); err != nil {
		c.JSON(http.StatusOK,Response{Code:400,Message:"删除失败"})
		return
	}

	if err := articleModel.Delete(articleIds); err != nil {
		c.JSON(http.StatusOK,Response{Code:400,Message:"删除失败"})
		return
	}

	c.JSON(http.StatusOK,Response{Code:200,Message:"删除成功"})
}

// 文章列表
func ArticleList(c *gin.Context)  {
	var articleModel models.Article

	pageIndex, _ := strconv.Atoi(c.DefaultQuery("page","1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("limit","20"))
	articleTitle := c.Query("article_title")
	articleStatus, _ := strconv.Atoi(c.DefaultQuery("article_status","-1"))
	tagId, _ := strconv.ParseInt(c.DefaultQuery("tag_id","-1"),10,64)

	_articles, total, err := articleModel.GetArticlePage(pageSize,pageIndex,articleTitle,articleStatus,tagId)
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