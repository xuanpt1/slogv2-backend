package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"slogv2/src/main/entity"
	"slogv2/src/main/service"
)

func GetOptionsList(c *gin.Context) {

	optionsList, status, err := service.GetOptionList()

	if !ResponseHandler(c, status, err, "配置列表获取成功", optionsList) {
		log.Println(fmt.Sprintf("发生未知错误: %s", err.Error()))
	}
}

func GetOptionsByName(c *gin.Context) {
	key := ParamHandler(c, "name")["name"]

	options, status, err := service.GetOptionByKey(key)

	if !ResponseHandler(c, status, err, "配置获取成功", options) {
		log.Println(fmt.Sprintf("发生未知错误: %s", err.Error()))
	}
}

func CreateOption(c *gin.Context) {
	var options entity.Options
	_ = c.ShouldBind(&options)

	status, err := service.CreateOption(&options)

	if !ResponseHandler(c, status, err, "配置创建成功", nil) {
		log.Println(fmt.Sprintf("发生未知错误: %s", err.Error()))
	}
}

func DeleteOption(c *gin.Context) {
	key := ParamHandler(c, "name")["name"]

	status, err := service.DeleteOption(key)

	if !ResponseHandler(c, status, err, "配置删除成功", nil) {
		log.Println(fmt.Sprintf("发生未知错误: %s", err.Error()))
	}
}

func UpdateOption(c *gin.Context) {
	var options entity.Options
	_ = c.ShouldBind(&options)

	status, err := service.UpdateOption(&options)

	if !ResponseHandler(c, status, err, "配置更新成功", nil) {
		log.Println(fmt.Sprintf("发生未知错误: %s", err.Error()))
	}
}
