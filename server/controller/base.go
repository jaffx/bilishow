package controller

import (
	"github.com/gin-gonic/gin"
)

type Controller interface {
	Before(*gin.Context)
	Main(*gin.Context)
	After(*gin.Context)
}

func NewHandleFunc(controller Controller) gin.HandlerFunc {
	return func(c *gin.Context) {
		controller.Before(c)
		controller.Main(c)
		controller.After(c)
	}
}
