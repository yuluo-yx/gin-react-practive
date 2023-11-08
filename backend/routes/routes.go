package routes

import (
	api "backend/api/v1"
	"backend/middleware"
	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {

	r := gin.Default()

	r.Use(middleware.Cors())
	v1 := r.Group("/api/v1")
	{
		// test ping request
		v1.GET("ping", func(c *gin.Context) {
			c.JSON(200, "success")
		})

		// 用户操作路由
		v1.POST("user/register", api.UserRegister)
		v1.POST("user/login", api.UserLogin)
		v1.GET("user/get", api.UserGet)

		// 留言操作路由
		v1.POST("msg/publish", api.MsgPublish)
		v1.GET("msg/get", api.MsgGet)
		v1.GET("msg/list", api.ListMsg)
		v1.DELETE("msg/delete", api.MsgDelete)
	}

	return r
}
