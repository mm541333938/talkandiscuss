package v1

import (
	"api/pkg/e"
	"github.com/gin-gonic/gin"
)

func GetAppVersionTest(c *gin.Context) {
	c.JSON(e.SUCCESS, gin.H{
		"Code": e.SUCCESS,
		"Msg":  e.GetMsg(e.SUCCESS),
		"Data": "返回数据成功",
	})
}
