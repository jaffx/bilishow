package controller

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type ControllerTest struct {
}

func (ctl *ControllerTest) Before(ctx *gin.Context) {

}

func (ctl *ControllerTest) Main(ctx *gin.Context) {
	ctx.String(http.StatusOK, "Hello World!")
	time.Sleep(time.Second)
}

func (ctl *ControllerTest) After(ctx *gin.Context) {

}
