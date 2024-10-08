package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"slogv2/src/main/utils/customError"
)

func ParamHandler(c *gin.Context, args ...string) map[string]string {
	mapArgs := make(map[string]string, len(args))
	if len(args) > 0 {
		for _, argName := range args {
			param := c.Param(argName)
			query := c.Query(argName)
			//log.Println(fmt.Sprintf("获取到参数: %s=%s", argName, param))
			//log.Println(fmt.Sprintf("获取到参数: %s=%s", argName, query))
			var _arg string
			if query != "" {
				_arg = query
			} else if param != "" {
				_arg = param
			} else {
				c.JSON(http.StatusBadRequest, gin.H{
					"code": http.StatusBadRequest,
					"msg":  "参数错误 未获取到对应参数" + argName + "   请检查参数是否正确",
				})
				return nil
			}
			mapArgs[argName] = _arg
		}
		//log.Println(fmt.Sprintf("获取到参数: %v", mapArgs))
		return mapArgs
	} else {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": http.StatusBadRequest,
			"msg":  "参数错误 参数列表为空",
		})
		return nil
	}
}

// TODO 重新规范Service类中https status code使用
func ResponseHandler(c *gin.Context, status int, err error, successMsg string, data interface{}) bool {
	if status == http.StatusOK {
		if err != nil {
			c.JSON(http.StatusOK, gin.H{
				"code": customError.GetCode(err),
				"msg":  err.Error(),
				"data": data,
			})
			return true
		}
		c.JSON(http.StatusOK, gin.H{
			"code": status,
			"msg":  successMsg,
			"data": data,
		})
		return true
	} else {
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"code": status,
				"msg":  err.Error(),
			})
			return false
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{
				"code": status,
				"msg":  customError.GetMsg(status),
			})
			return false
		}
	}
}
