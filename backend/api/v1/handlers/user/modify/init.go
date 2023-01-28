package modify

import "github.com/gin-gonic/gin"

func LoadModify(path *gin.RouterGroup) {
	subpath := path.Group("/modify")

	subpath.POST("/", modify)
}
