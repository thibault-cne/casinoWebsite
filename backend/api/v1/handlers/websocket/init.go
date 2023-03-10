package websocket

import (
	"net/http"

	"casino.website/api/v1/middlewares"
	"github.com/gin-gonic/gin"
	"github.com/yyewolf/gosf"
)

var (
	Ws     http.Handler = gosf.GetHandler(map[string]interface{}{})
	Engine *gin.Engine
)

func LoadWebsocket(path *gin.RouterGroup, engine *gin.Engine) {
	Engine = engine

	subpath := path.Group("/ws")

	subpath.GET("/", middlewares.AuthRequired(), Wrapper())
}
