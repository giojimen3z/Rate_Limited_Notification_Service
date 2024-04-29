package health

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

const (
	Pong string = "pong"
)

var (
	PingController pingControllerInterface = &pingController{}
)

type pingControllerInterface interface {
	Ping(c *gin.Context)
}

type pingController struct{}

func (controller *pingController) Ping(c *gin.Context) {
	c.String(http.StatusOK, Pong)
}
