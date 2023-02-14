package auth

import (
	"regexp"
	"time"

	"casino.website/internal/env"
	"casino.website/internal/models"
	"casino.website/internal/utils"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

// registerForm is the struct that will be used to bind the request body
type registerForm struct {
	Username       string `json:"username" binding:"required"`
	Password       string `json:"password" binding:"required"`
	PasswordVerify string `json:"passwordVerify" binding:"required"`
	Code           string `json:"code" binding:"required"`
}

var (
	usernameRegex = regexp.MustCompile(`^[a-zA-Z0-9_]{3,16}$`)
	passwordRegex = regexp.MustCompile(`^[a-zA-Z0-9!@#$%^&*()_+-=]{8,32}$`)
)

func Register(c *gin.Context) {
	var f registerForm
	// Bind the request body to the struct
	if err := c.MustBindWith(&f, binding.JSON); err != nil {
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

	// Check if the code is valid
	claim, err := models.GetClaim(f.Code)

	if err != nil {
		c.JSON(400, gin.H{
			"success": false,
			"status":  "Le code entrée n'existe pas",
		})
		return
	}

	if claim.Use <= 0 {
		c.JSON(400, gin.H{
			"success": false,
			"status":  "Le code utilisée n'est plus actifs",
		})
		return
	}

	// Check if the username is already taken
	if _, err := models.GetUserByUsername(f.Username); err != nil {
		c.JSON(400, gin.H{
			"success": false,
			"status":  "Nom d'utilisateur déjà pris",
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
		ID:        utils.GenerateUserId(),
		Wallet:    0,
		Status:    "user",
		Username:  f.Username,
		Password:  pHash,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	if claim.Code == env.Config.SuperAdminClaim {
		user.Status = "super admin"
	}

	// Update the code used
	claim.Use -= 1
	err = claim.Save()

	if err != nil {
		c.JSON(500, gin.H{
			"success": false,
			"status":  "Erreur interne",
		})
		return
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
