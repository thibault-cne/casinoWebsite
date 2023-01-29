package get

import "github.com/gin-gonic/gin"

func LoadGet(path *gin.RouterGroup) {
	subpath := path.Group("/get")

	subpath.GET("/wallet", wallet)
	subpath.GET("/all", all)
	subpath.GET("/picture/:userId", getPicture)
}
