
package models

import (
	orm "nuxt-blog-api/db"
	"nuxt-blog-api/utils"
	"log"
)

type ArticleTag struct {
	ArticleID  int64  `json:"article_id"`
	TagID      int64  `json:"tag_id"`
	CreateTime string `son:"create_time"`
}

// 创建文章标签
func (at ArticleTag) Create() (err error) {
	at.CreateTime = utils.GetCurrntTime()
	if err = orm.Eloquent.Table("article_tags").Create(&at).Error; err != nil {
		log.Println(err.Error())
	}
	return
}

// 删除文章标签
func (at ArticleTag) Delete(articleId int64,tagIds []int64)(err error)  {
	if err = orm.Eloquent.Table("article_tags").Where("article_id = ?",articleId).Where("tag_id in (?)",tagIds).Delete(&at).Error; err != nil {
		log.Println(err.Error())
	}
	return
}