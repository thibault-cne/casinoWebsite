package games

import "github.com/gin-gonic/gin"

func LoadAdminGames(path *gin.RouterGroup) {
	subpath := path.Group("/games")

	subpath.GET("/roulette/toggle", toggleRoulette)
}
