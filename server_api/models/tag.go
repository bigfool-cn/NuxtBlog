package models

import (
	"github.com/jinzhu/gorm"
	orm "nuxt-blog-api/db"
	"log"
	"nuxt-blog-api/utils"
)

type Tag struct {
	TagID       int64    `gorm:"primary_key;AUTO_INCREMENT" json:"tag_id"`
	TagName     string	`gorm:"size:20" json:"tag_name" binding:"required,max=20"`
	TagIcon     string	`json:"tag_icon" binding:"required"`
	UpdateTime  string  `gorm:"default:NULL" json:"update_time"`
	CreateTime  string  `json:"create_time"`
}

// 获取标签
func (t Tag) GetTag() (tag Tag, err error)  {
	if err = orm.Eloquent.Table("tags").Where(&t).First(&tag).Error; err != nil {
		log.Println(err.Error())
		if err == gorm.ErrRecordNotFound  {
			err = nil
		}
	}
	return
}

// 获取标签
func (t Tag) GetTags() (tags []Tag, err error)  {
	if err = orm.Eloquent.Table("tags").Where(&t).Find(&tags).Error; err != nil {
		log.Println(err.Error())
		if err == gorm.ErrRecordNotFound  {
			err = nil
		}
	}
	return
}


// 创建标签
func (t Tag) Create() (tagId int64, err error)  {
	t.CreateTime = utils.GetCurrntTime()
	if err = orm.Eloquent.Table("tags").Create(&t).Error; err != nil {
		log.Println(err.Error())
	}
	tagId = t.TagID
	return
}

// 修改标签
func (t Tag) Update() (err error) {
	t.UpdateTime = utils.GetCurrntTime()
	if err = orm.Eloquent.Table("tags").Omit("create_time").Save(&t).Error; err != nil {
		log.Println(err.Error())
	}
	return
}


// 删除标签
func (t Tag) Delete(tagIds []int64)(err error)  {
	tran := orm.Eloquent.Begin()

	if err = orm.Eloquent.Table("tags").Where("tag_id in (?)",tagIds).Delete(&t).Error; err != nil {
		tran.Rollback()
		log.Println(err.Error())
		return
	}

	if err = orm.Eloquent.Table("article_tags").Where("tag_id in (?)",tagIds).Delete(&t).Error; err != nil {
		tran.Rollback()
		log.Println(err.Error())
		return
	}

	err = tran.Commit().Error
	return
}

// 获取列表
func (t Tag) GetTagPage(pageSize int, pageIndex int, tagName string) (Tags []Tag,count int64,err error)  {
	table := orm.Eloquent.Table("tags")
	if tagName != "" {
		table = table.Where("tags.tag_name = ?",tagName)
	}
	if err = table.Offset((pageIndex -1) * pageSize).Limit(pageSize).Order("tags.create_time desc").Find(&Tags).Error; err != nil {
		log.Println(err.Error())
		if err == gorm.ErrRecordNotFound  {
			err = nil
		}
	}
	table.Count(&count)
	return
}
