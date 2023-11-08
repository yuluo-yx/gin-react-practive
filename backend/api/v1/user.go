package v1

import (
	"backend/pkg/utils"
	"backend/serializer"
	"backend/services"
	"github.com/gin-gonic/gin"
	"net/http"
)

func UserRegister(c *gin.Context) {

	var userRegister services.UserService

	if err := c.ShouldBind(&userRegister); err == nil {
		res := userRegister.Register(c.Request.Context())
		c.JSON(http.StatusOK, res)
	} else {
		// 加入日志处理
		utils.LogrusObj.Info("user Register err: ", err)
		c.JSON(http.StatusBadRequest, serializer.ErrorResponse(err))
	}
}

func UserLogin(c *gin.Context) {

	var userLogin services.UserService

	if err := c.ShouldBind(&userLogin); err == nil {
		res := userLogin.Login(c.Request.Context())
		c.JSON(http.StatusOK, res)
	} else {
		utils.LogrusObj.Info("user Login err: ", err)
		c.JSON(http.StatusBadRequest, serializer.ErrorResponse(err))
	}
}

func UserGet(c *gin.Context) {

	var userInfo services.UserService

	if err := c.ShouldBind(&userInfo); err == nil {
		res := userInfo.Get(c.Request.Context())
		c.JSON(http.StatusOK, res)
	} else {
		utils.LogrusObj.Info("get user info err: ", err)
		c.JSON(http.StatusBadRequest, serializer.ErrorResponse(err))
	}

}
