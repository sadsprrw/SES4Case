package handlers

import (
	"example/SES4Case/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Subscribe(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		email := c.PostForm("email")

		if email == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Email is required"})
			return
		}

		var existing models.Email
		result := db.First(&existing, "address = ?", email)

		if result.Error == nil {
			c.JSON(http.StatusConflict, gin.H{"error": "Email already subscribed"})
			return
		}

		if result.Error != nil && result.Error != gorm.ErrRecordNotFound {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Database error"})
			return
		}

		newEmail := models.Email{Address: email}
		db.Create(&newEmail)
		c.JSON(http.StatusOK, gin.H{"message": "Email subscribed"})
	}
}
