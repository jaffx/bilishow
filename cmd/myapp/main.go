package main

import (
	"bilishow"

	"github.com/gin-gonic/gin"
)

func main() {
	engine := gin.Default()
	bilishow.RegisterView(engine)
	engine.Run(":8080")
}
