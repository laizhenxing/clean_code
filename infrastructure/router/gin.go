package router

import (
	"time"

	"github.com/gin-gonic/gin"

	middleware "cc/infrastructure/middleware/mgin"
)

var Router *gin.Engine

func InitGinRouter() *gin.Engine {
	Router = gin.New()

	Router.Use(gin.Logger())
	Router.Use(gin.Recovery())

	AddRouters()

	return Router
}

func AddRouters() {
	v1 := Router.Group("v1")
	{
		appCtrl := getAppGinHandler()
		//appCtrl := getAppMongoHandler()
		v1.POST("/app", appCtrl.Add())
		v1.POST("/app/:id", appCtrl.Update())
		// 限制每秒100个请求,局部中间件 RateLimit
		v1.GET("/apps", middleware.RateLimit(time.Second, 100), appCtrl.FindAll())
		v1.GET("/app/:id", appCtrl.Get())
		v1.DELETE("/app/:id", appCtrl.Delete())

		// user
		userCtrl := getUserHandler()
		v1.POST("/user", userCtrl.Create())
		v1.Use(middleware.JWT()).POST("/user/:id", userCtrl.Update())
		v1.GET("/users", userCtrl.FindAll())
		v1.GET("/user/:id", userCtrl.FindOne())
		v1.DELETE("/user/:id", userCtrl.Delete())
	}
}
