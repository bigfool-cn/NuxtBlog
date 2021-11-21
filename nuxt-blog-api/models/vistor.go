package models

import (
	"github.com/jinzhu/gorm"
	orm "nuxt-blog-api/db"
	"log"
	"nuxt-blog-api/utils"
)

type Vistor struct {
	VistorID       int64    `gorm:"primary_key;AUTO_INCREMENT" json:"vistor_id"`
	VistorIP       string	`gorm:"size:20" json:"vistor_ip" binding:"required"`
	UserAgent      string	`json:"user_agent" binding:"required"`
	VistorPath     string	`json:"vistor_path" binding:"required"`
	CreateTime     string   `json:"create_time"`
}

// 获取访客记录
func (v Vistor) GetVistor() (vistor Vistor, err error)  {
	if err = orm.Eloquent.Table("vistors").Where(&v).First(&vistor).Error; err != nil {
		log.Println(err.Error())
		if err == gorm.ErrRecordNotFound  {
			err = nil
		}
	}
	return
}

// 创建访客记录
func (v Vistor) Create() (vistorId int64, err error)  {
	v.CreateTime = utils.GetCurrntTime()
	if err = orm.Eloquent.Table("vistors").Create(&v).Error; err != nil {
		log.Println(err.Error())
	}
	vistorId = v.VistorID
	return
}


// 删除访客记录
func (v Vistor) Delete(vistorIds []int64)(err error)  {

	if err = orm.Eloquent.Table("vistors").Where("vistor_id in (?)",vistorIds).Delete(&v).Error; err != nil {
		log.Println(err.Error())
	}

	return
}

// 获取列表
func (v Vistor) GetVistorPage(pageSize int, pageIndex int, vistorIP string, startTime string, endTime string) (Vistors []Vistor,count int64,err error)  {
	table := orm.Eloquent.Table("vistors")

	if vistorIP != "" {
		table = table.Where("vistor_ip = ?",vistorIP)
	}

	if startTime != "" {
		table = table.Where("create_time >= ?",startTime)
	}

	if endTime != "" {
		table = table.Where("create_time <= ?",endTime)
	}

	if err = table.Offset((pageIndex -1) * pageSize).Limit(pageSize).Order("vistors.create_time desc").Find(&Vistors).Error; err != nil {
		log.Println(err.Error())
		if err == gorm.ErrRecordNotFound  {
			err = nil
		}
	}

	table.Count(&count)
	return
}
