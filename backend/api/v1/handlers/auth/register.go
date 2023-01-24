package auth

import (
	"regexp"
	"time"

	"casino.website/internal/models"
	"casino.website/internal/utils"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

// registerForm is the struct that will be used to bind the request body
type registerForm struct {
	Username       string `json:"username" binding:"required"`
	Email          string `json:"email" binding:"required"`
	Password       string `json:"password" binding:"required"`
	PasswordVerify string `json:"passwordVerify" binding:"required"`
}

var (
	usernameRegex = regexp.MustCompile(`^[a-zA-Z0-9_]{3,16}$`)
	emailRegex    = regexp.MustCompile(`^[a-zA-Z0-9_.+-]+@[a-zA-Z0-9-]+\.[a-zA-Z0-9-.]+$`)
	passwordRegex = regexp.MustCompile(`^[a-zA-Z0-9!@#$%^&*()_+-=]{8,32}$`)
)

func Register(c *gin.Context) {
	var f registerForm
	// Bind the request body to the struct
	if err := c.ShouldBindBodyWith(&f, binding.JSON); err != nil {
		c.JSON(400, gin.H{
			"success": false,
			"status":  "Les informations fournies ne sont pas valides",
		})
		return
	}

	// Check if the passwords match
	if f.Password != f.PasswordVerify {
		c.JSON(400, gin.H{
			"success": false,
			"status":  "Les mots de passe ne correspondent pas",
		})
		return
	}

	// Check if the username is already taken
	if _, err := models.GetUserByUsername(f.Username); err == nil {
		c.JSON(400, gin.H{
			"success": false,
			"status":  "Nom d'utilisateur déjà pris",
		})
		return
	}

	// Check if the email is already taken
	if _, err := models.GetUserByEmail(f.Email); err == nil {
		c.JSON(400, gin.H{
			"success": false,
			"status":  "Email déjà utilisée",
		})
		return
	}

	// Verify the fields with regex
	if !usernameRegex.MatchString(f.Username) {
		c.JSON(400, gin.H{
			"success": false,
			"status":  "Nom d'utilisateur invalide, doit être compris entre 3 et 16 caractères et ne contenir que des lettres, des chiffres et des underscores",
		})
		return
	}

	if !emailRegex.MatchString(f.Email) {
		c.JSON(400, gin.H{
			"success": false,
			"status":  "Email invalide, doit être une adresse email valide",
		})
		return
	}

	if !passwordRegex.MatchString(f.Password) {
		c.JSON(400, gin.H{
			"success": false,
			"status":  "Mot de passe invalide, doit être compris entre 8 et 32 caractères et ne contenir que des lettres, des chiffres et des caractères spéciaux",
		})
		return
	}

	// Hash the password
	pHash, err := models.HashPassword(f.Password)
	if err != nil {
		c.JSON(500, gin.H{
			"success": false,
			"status":  "Erreur interne",
		})
		return
	}

	// Create the user
	user := models.User{
		ID:        utils.Generate(),
		Username:  f.Username,
		Email:     f.Email,
		Password:  pHash,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	// Save the user in the database
	if err := user.Create(); err != nil {
		c.JSON(500, gin.H{
			"success": false,
			"status":  "Erreur interne",
		})
		return
	}

	// Return positive response
	c.JSON(200, gin.H{
		"success": true,
		"status":  "Compte créé avec succès",
	})
}
