package models

import (
	"github.com/jinzhu/gorm"
	"github.com/sirupsen/logrus"
	orm "nuxt-blog-api/db"
	"nuxt-blog-api/utils"
)

type Article struct {
	ArticleID       int64   `gorm:"primary_key;AUTO_INCREMENT" json:"article_id"`
	ArticleTitle    string  `gorm:"size:500" json:"article_title" binding:"required,len=500"`
	ArticleAuthor   string  `gorm:"default:'bigfool'" json:"article_author" binding:""`
	ArticleRead     int64   `json:"article_read"`
	ArticleStatus   int     `gorm:"default:1" json:"article_status" binding:""`
	ArticleDesc     string  `gorm:"size:500" json:"article_desc" binding:""`
	ArticleContent  string  `json:"article_content" binding:"required"`
	UpdateTime      string  `gorm:"default:NULL" json:"update_time"`
	CreateTime      string  `json:"create_time"`

	Tags            []Tag `gorm:"association_autoupdate:false;many2many:article_tags;foreignkey:article_id;association_foreignkey:tag_id;association_jointable_foreignkey:tag_id;jointable_foreignkey:article_id;" json:"tags"`
}

type ArticleForm struct {
	ArticleTitle    string  `json:"article_title" binding:"required,max=500"`
	ArticleStatus   int     `json:"article_status" binding:""`
	ArticleDesc     string  `json:"article_desc" binding:"max=500"`
	ArticleContent  string  `json:"article_content" binding:"required"`
	TagIds          []int64 `json:"tag_ids" binding:"required"`
}

var lgr *logrus.Logger

func init()  {
	lgr = utils.DefaultLogger(false)
}

// 获取文章
func (a Article) GetArticle() (article Article, err error)  {
	if err = orm.Eloquent.Table("articles").Preload("Tags").Where(&a).First(&article).Error; err != nil {
		lgr.Error(err.Error())
		if err == gorm.ErrRecordNotFound  {
			err = nil
		}
	}
	return
}

// 根据id 获取文章
func (a Article) GetArticleById(articleId int64) (article Article, err error)  {
	if err = orm.Eloquent.Table("articles").Preload("Tags").Take(&article,articleId).Error; err != nil {
		lgr.Error(err.Error())
		if err == gorm.ErrRecordNotFound  {
			err = nil
		}
	}
	return
}

// 创建接口
func (a Article) Create() (err error)  {
	a.CreateTime = utils.GetCurrntTime()
	if err = orm.Eloquent.Table("articles").Create(&a).Error; err != nil {
		lgr.Error(err.Error())
	}
	return
}

// 修改文章
func (a Article) Update() (err error) {
	var (
		oldTagIds       []int64
		newTagIds       []int64
		articleTagModel ArticleTag
		article         Article
	)

	a.UpdateTime = utils.GetCurrntTime()
	tran := orm.Eloquent.Begin()

	if article, err = a.GetArticleById(a.ArticleID); err != nil {
		lgr.Error(err.Error())
		tran.Rollback()
		return
	}

	for _, t := range article.Tags {
		oldTagIds = append(oldTagIds,t.TagID)
	}
	for _, t := range a.Tags {
		newTagIds = append(newTagIds,t.TagID)
	}

	// 删除取消的标签
	delTagIds, err := utils.Difference(oldTagIds,newTagIds)
	if err != nil {
		lgr.Error(err.Error())
		tran.Rollback()
		return err
	}
	if len(delTagIds.([]int64)) > 0 {
		if err := articleTagModel.Delete(a.ArticleID, delTagIds.([]int64)); err != nil {
			lgr.Error(err.Error())
			tran.Rollback()
			return err
		}
	}

	// 添加新的标签
	addTagIds, err := utils.Difference(newTagIds,oldTagIds)
	if err != nil {
		lgr.Error(err.Error())
		tran.Rollback()
		return err
	}
	if len(addTagIds.([]int64)) > 0 {
		for _, tagId := range addTagIds.([] int64) {
			ArticleTag := ArticleTag{
				ArticleID:     a.ArticleID,
				TagID:         tagId,
			}
			if err := ArticleTag.Create(); err != nil {
				lgr.Error(err.Error())
				tran.Rollback()
				return err
			}
		}
	}

	if err = orm.Eloquent.Table("articles").Omit("create_time").Save(&a).Error; err != nil {
		lgr.Error(err.Error())
		tran.Rollback()
		return
	}

	err = tran.Commit().Error
	return
}

// 更新阅读量
func (a Article) UpdateArticleRead(articleId int64, read int64) (err error)  {
	if err = orm.Eloquent.Table("articles").Where("article_id = ?",articleId).Update("article_read", read).Error; err != nil {
		lgr.Error(err.Error())
	}
	return
}


// 删除文章
func (a Article) Delete(articleIds []int64)(err error)  {
	tran := orm.Eloquent.Begin()
	if err = orm.Eloquent.Table("articles").Where("article_id in (?)",articleIds).Delete(&a).Error; err != nil {
		tran.Rollback()
		lgr.Error(err.Error())
		return
	}

	if err = orm.Eloquent.Table("article_tags").Where("article_id in (?)",articleIds).Delete(&a).Error; err != nil {
		tran.Rollback()
		lgr.Error(err.Error())
		return
	}
	err = tran.Commit().Error
	return
}

// 获取列表
func (a Article) GetArticlePage(pageSize int, pageIndex int, articleTitle string, articleStatus int, tagId int64) (Articles []Article,count int64,err error)  {
	table := orm.Eloquent.Table("articles")
	if articleTitle != "" {
		table = table.Where("articles.article_title like ?",articleTitle)
	}
	if articleStatus != -1 {
		table = table.Where("articles.article_status = ?",articleStatus)
	}

	if tagId != -1 {
		table = table.Where("article_id IN (SELECT article_id FROM article_tags WHERE tag_id = ?)",tagId)
	}
	table = table.Preload("Tags")

	table = table.Select([]string{"article_id","article_author","article_title","article_desc","article_status","article_read","update_time","create_time"})

	if err = table.Offset((pageIndex -1) * pageSize).Limit(pageSize).Order("articles.create_time desc").Find(&Articles).Error; err != nil {
		lgr.Error(err.Error())
		if err == gorm.ErrRecordNotFound  {
			err = nil
		}
	}
	table.Count(&count)
	return
}