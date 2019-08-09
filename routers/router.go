package routers

import (
	"github.com/gin-gonic/gin"
	"web_login/controller"
)


func InitRouter() *gin.Engine {

	router := gin.Default()
	router.LoadHTMLGlob("view/*")
	router.Static("/static", "./static")

	//注册：
	router.GET("/register", controller.RegisterGet)
	router.POST("/register", controller.RegisterPost)

	//登录
	router.GET("/login", controller.LoginGet)
	router.POST("/login", controller.LoginPost)

	//强制修改密码
	router.GET("/change_password", controller.ChangepwGet)
	router.POST("/change_password", controller.ChangepwPost)

		//验证
		router.Use(controller.JWTAuth())

		{
			//获取homepage
			router.GET("/", controller.HomePage)

			//进入db_info并调取对应功能
			router.GET("/db_info", controller.DB_Info_Get)
			router.POST ("/db_info", controller.DB_Info_Post)
		}
	return router
}
