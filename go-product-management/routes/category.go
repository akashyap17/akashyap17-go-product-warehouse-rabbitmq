package routes

import (
	"github.com/akashyap17/go-product-management/db"
	"github.com/akashyap17/go-product-management/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func CreateCategory(c *gin.Context) {
	var category models.Category
	if err := c.ShouldBindJSON(&category); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := db.DB.Create(&category).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Category created successfully"})
}

func GetCategories(c *gin.Context) {
	var categories []models.Category
	if err := db.DB.Find(&categories).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, categories)
}

// Similarly, you can implement UpdateCategory and DeleteCategory...
