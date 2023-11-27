package v1

import (
	"backend/pkg/utils"
	"backend/serializer"
	"backend/services"
	"github.com/gin-gonic/gin"
	"net/http"
)

func MsgPublish(c *gin.Context) {
	var msgService services.MsgService

	if err := c.ShouldBind(&msgService); err == nil {
		res := msgService.Publish(c.Request.Context())
		c.JSON(http.StatusOK, res)
	} else {
		// 加入日志处理
		utils.LogrusObj.Info("Msg publish err: ", err)
		c.JSON(http.StatusBadRequest, serializer.ErrorResponse(err))
	}
}

func MsgDelete(c *gin.Context) {
	var msgService services.MsgService

	if err := c.ShouldBind(&msgService); err == nil {
		res := msgService.Delete(c.Request.Context())
		c.JSON(http.StatusOK, res)
	} else {
		// 加入日志处理
		utils.LogrusObj.Info("Msg publish err: ", err)
		c.JSON(http.StatusBadRequest, serializer.ErrorResponse(err))
	}
}

//func MsgGet(c *gin.Context) {
//	var msgService services.MsgService
//
//	if err := c.ShouldBind(&msgService); err == nil {
//		res := msgService.Get(c.Request.Context())
//		c.JSON(http.StatusOK, res)
//	} else {
//		// 加入日志处理
//		utils.LogrusObj.Info("Msg get err: ", err)
//		c.JSON(http.StatusBadRequest, serializer.ErrorResponse(err))
//	}
//}

func ListMsg(c *gin.Context) {
	var msgService services.MsgService

	if err := c.ShouldBind(&msgService); err == nil {
		res := msgService.ListMsg(c.Request.Context())
		c.JSON(http.StatusOK, res)
	} else {
		// 加入日志处理
		utils.LogrusObj.Info("Msg get err: ", err)
		c.JSON(http.StatusBadRequest, serializer.ErrorResponse(err))
	}
}
