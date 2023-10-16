package bilishow

import (
	"bilishow/server/controller"

	"github.com/gin-gonic/gin"
)

func RegisterView(engine *gin.Engine) {
	// engine.Use(midware.RunTimeMiddleWare())
	engine.Static("/static", "./web/static")
	engine.StaticFile("/", "./web/static/html/index.html")
	engine.GET("/test", controller.NewHandleFunc(&controller.ControllerTest{}))
	testGroup := engine.Group("/test")
	{
		testGroup.GET("/test", controller.NewHandleFunc(&controller.ControllerTest{}))
	}
}
