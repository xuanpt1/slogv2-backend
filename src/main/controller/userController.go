package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"slogv2/src/main/entity"
	"slogv2/src/main/service"
)

func GetUserByUid(c *gin.Context) {
	uid := ParamHandler(c, "uid")["uid"]

	user, status, err := service.GetUserByUid(uid)

	if !ResponseHandler(c, status, err, "用户信息查询成功", user) {
		log.Println(fmt.Sprintf("发生未知错误: %s", err.Error()))
	}

}

func GetUserList(c *gin.Context) {
	userList, status, err := service.GetUserList()

	if !ResponseHandler(c, status, err, "用户列表查询成功", userList) {
		log.Println(fmt.Sprintf("发生未知错误: %s", err.Error()))
	}
}

func GetUserByUsername(c *gin.Context) {
	username := ParamHandler(c, "username")["username"]

	user, status, err := service.GetUserByUsername(username)

	if !ResponseHandler(c, status, err, "用户信息查询成功", user) {
		log.Println(fmt.Sprintf("发生未知错误: %s", err.Error()))
	}
}

func CreateUser(c *gin.Context) {
	var user entity.User
	_ = c.ShouldBind(&user)

	status, err := service.CreateUser(&user)

	if !ResponseHandler(c, status, err, "用户创建成功", user) {
		log.Println(fmt.Sprintf("发生未知错误: %s", err.Error()))
	}
}

func UpdateUser(c *gin.Context) {
	var user entity.User
	_ = c.ShouldBind(&user)

	status, err := service.UpdateUser(&user)

	if !ResponseHandler(c, status, err, "用户信息更新成功", user) {
		log.Println(fmt.Sprintf("发生未知错误: %s", err.Error()))
	}
}

func DeleteUser(c *gin.Context) {
	uid := ParamHandler(c, "uid")["uid"]

	status, err := service.DeleteUser(&entity.User{Uid: uid})

	if !ResponseHandler(c, status, err, "用户删除成功", uid) {
		log.Println(fmt.Sprintf("发生未知错误: %s", err.Error()))
	}
}
