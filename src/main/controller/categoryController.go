package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"slogv2/src/main/entity"
	"slogv2/src/main/service"
	"strconv"
)

func CreateCategory(c *gin.Context) {
	var category entity.Category
	_ = c.ShouldBind(&category)

	status, err := service.CreateCategory(&category)

	if !ResponseHandler(c, status, err, "分类创建成功", nil) {
		log.Println(fmt.Sprintf("发生未知错误: %s", err.Error()))
	}
}

func GetCategoryList(c *gin.Context) {
	categoryList, status, err := service.GetCategoryList()

	if !ResponseHandler(c, status, err, "分类列表获取成功", categoryList) {
		log.Println(fmt.Sprintf("发生未知错误: %s", err.Error()))
	}
}

func GetCategory(c *gin.Context) {
	cid := ParamHandler(c, "cid")["cid"]
	_cid, _ := strconv.Atoi(cid)

	category, status, err := service.GetCategoryByCid(_cid)

	if !ResponseHandler(c, status, err, "分类获取成功", category) {
		log.Println(fmt.Sprintf("发生未知错误: %s", err.Error()))
	}
}

func UpdateCategory(c *gin.Context) {
	var category entity.Category
	_ = c.ShouldBind(&category)

	status, err := service.UpdateCategory(&category)

	if !ResponseHandler(c, status, err, "分类更新成功", nil) {
		log.Println(fmt.Sprintf("发生未知错误: %s", err.Error()))
	}
}

func DeleteCategory(c *gin.Context) {
	cid := ParamHandler(c, "cid")["cid"]

	status, err := service.DeleteCategoryById(cid)

	if !ResponseHandler(c, status, err, "分类删除成功", nil) {
		log.Println(fmt.Sprintf("发生未知错误: %s", err.Error()))
	}
}
