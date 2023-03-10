package claims

import "github.com/gin-gonic/gin"

func LoadClaims(path *gin.RouterGroup) {
	subpath := path.Group("/claims")

	subpath.GET("/get", getAllClaims)
	subpath.POST("/create", newClaim)
	subpath.DELETE("/delete", deleteClaim)
}
