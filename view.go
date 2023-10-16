package bilishow

import (
	"bilishow/server/controller"
	"bilishow/server/midware"

	"github.com/gin-gonic/gin"
)

func RegisterView(engine *gin.Engine) {
	engine.Use(midware.RunTimeMiddleWare())
	engine.Static("/static", "./web/static")
	engine.StaticFile("/", "./web/static/html/index.html")
	testGroup := engine.Group("/test")
	{
		testGroup.GET("/test", controller.NewHandleFunc(&controller.ControllerTest{}))
	}
}
