package auth

import (
	"fmt"
	"net/http"

	"casino.website/internal/models"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

// loginForm is the struct that will be used to bind the request body
type loginForm struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func Login(c *gin.Context) {
	var f loginForm
	// Bind the request body to the struct
	if err := c.ShouldBindBodyWith(&f, binding.JSON); err != nil {
		c.JSON(400, gin.H{
			"success": false,
			"status":  "Les informations fournies ne sont pas valides",
		})
		return
	}
	fmt.Println("Error 1")

	// Check if the user exists by username or email
	user, err := models.GetUserByUsername(f.Username)
	if err != nil {

		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"status":  "Nom d'utilisateur ou mot de passe incorrect",
		})
		return
	}
	fmt.Println("Error 2")
	fmt.Printf("%s %+v\n", f.Password, user)

	// Check if the password is correct
	if !user.CheckPasswordHash(f.Password) {
		c.JSON(400, gin.H{
			"success": false,
			"status":  "Nom d'utilisateur ou mot de passe incorrect",
		})
		return
	}
	fmt.Println("Error 3")

	// Store the user ID in the session
	session := sessions.Default(c)
	session.Set("user", user.ID)
	session.Save()
	fmt.Println(session.ID())

	c.JSON(200, gin.H{
		"success": true,
		"status":  "Vous êtes maintenant connecté",
		"user":    user,
	})
}
