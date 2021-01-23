package apis

import (
	"github.com/gin-gonic/gin"
	"log"
	"nuxt-blog-api/cache"
	"nuxt-blog-api/models"
	"strconv"
	"net/http"
)

// 获取所有标签
func GetTags (c *gin.Context){
	var tagModel models.Tag

	tagModel.TagName = c.Query("tag_name")
	tags, err := tagModel.GetTags()

	if err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusOK,Response{Code:400,Message:"获取数据失败"})
		return
	}

	c.JSON(http.StatusOK,Response{Code:200,Message:"获取成功",Data:tags})
}

// 创建标签
func CreateTag (c *gin.Context){
	var (
		tagForm models.Tag
		tagModel models.Tag
	)

	if err := c.BindJSON(&tagForm); err != nil {
		c.JSON(http.StatusOK,Response{Code:400,Message:"数据验证失败," + err.Error()})
		return
	}

	tagModel.TagName = tagForm.TagName

	if tag, _ := tagModel.GetTag(); tag.TagID > 0 {
		c.JSON(http.StatusOK,Response{Code:400,Message:"该标签已存在"})
		return
	}

	if _,err := tagForm.Create(); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusOK,Response{Code:400,Message:"添加失败"})
		return
	}

	cache.RedisCache.Del("blog_tags")

	c.JSON(http.StatusOK,Response{Code:200,Message:"添加成功"})
}

// 更新标签
func UpdateTag (c *gin.Context){
	var (
		tagForm models.Tag
		tagModel models.Tag
	)

	if err := c.BindJSON(&tagForm); err != nil {
		c.JSON(http.StatusOK,Response{Code:400,Message:"数据验证失败," + err.Error()})
		return
	}


	tagId, err := strconv.ParseInt(c.Param("tag_id"),10,64)
	if err != nil {
		c.JSON(http.StatusOK,Response{Code:400,Message:"参数验证失败"})
		return
	}

	tagModel.TagName = tagForm.TagName
	
	if tag, _ := tagModel.GetTag(); tag.TagID > 0 && tag.TagID != tagId {
		c.JSON(http.StatusOK,Response{Code:400,Message:"该标签已存在"})
		return
	}

	tagModel.TagID = tagId

	if err := tagForm.Update(); err != nil {
		c.JSON(http.StatusOK,Response{Code:400,Message:"修改失败"})
		return
	}

	cache.RedisCache.Del("blog_tags")

	c.JSON(http.StatusOK,Response{Code:200,Message:"修改成功"})
}

// 删除标签
func DeleteTag(c *gin.Context)  {
	var tagIds []int64
	if err := c.BindJSON(&tagIds); err != nil {
		c.JSON(http.StatusOK,Response{Code:400,Message:"删除失败"})
		return
	}

	var (
		tagModel models.Tag
	)
	if err := tagModel.Delete(tagIds); err != nil {
		c.JSON(http.StatusOK,Response{Code:400,Message:"删除失败"})
		return
	}

	cache.RedisCache.Del("blog_tags")

	c.JSON(http.StatusOK,Response{Code:200,Message:"删除成功"})
}
