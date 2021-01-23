package apis

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"nuxt-blog-api/models"
	"strconv"
)

// 删除访客记录
func DeleteVistor(c *gin.Context)  {
	var (
		vistorIds []int64
		vistorModel models.Vistor
	)

	if err := c.BindJSON(&vistorIds); err != nil {
		c.JSON(http.StatusOK,Response{Code:400,Message:"删除失败"})
		return
	}

	if err := vistorModel.Delete(vistorIds); err != nil {
		c.JSON(http.StatusOK,Response{Code:400,Message:"删除失败"})
		return
	}

	c.JSON(http.StatusOK,Response{Code:200,Message:"删除成功"})
}

// 访客列表
func VistorList(c *gin.Context)  {
	var vistorModel models.Vistor

	pageIndex, _ := strconv.Atoi(c.DefaultQuery("page","1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("limit","20"))
	vistorIp := c.Query("vistor_ip")
	startTime := c.Query("start_time")
	endTime := c.Query("end_time")
	log.Println("start_time",startTime)
	log.Println("endTime",endTime)

	_vistors, total, err := vistorModel.GetVistorPage(pageSize,pageIndex,vistorIp,startTime,endTime)
	if err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusOK,Response{Code:400,Message:"获取数据失败"})
		return
	}
	type vistors struct {
		Vistors  []models.Vistor   `json:"vistors"`
		Total     int64            `json:"total"`
	}
	c.JSON(http.StatusOK,Response{Code:200,Message:"获取成功",Data:&vistors{Vistors:_vistors,Total:total}})
}