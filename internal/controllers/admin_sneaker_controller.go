package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/muhammedfazall/go-ecommerce/internal/database"
	"github.com/muhammedfazall/go-ecommerce/internal/models"
	"gorm.io/gorm"
)

func GetAllSneakers(c *gin.Context) {
	var sneakers []models.Sneaker

	if err := database.DB.Order("created_at desc").Find(&sneakers).Error; err != nil {
		c.HTML(500, "list.html", gin.H{
			"Error": "failed to fetch sneakers",
		})
		return
	}
	c.HTML(http.StatusOK, "list.html", gin.H{
		"Sneakers": sneakers,
	})
}

func AddSneaker(c *gin.Context) {
	var input models.Sneaker

	if err := c.ShouldBindJSON(&input); err != nil {
		c.HTML(http.StatusBadRequest, "list.html", gin.H{
			"Error": "invalid input",
		})
		return
	}

	if err := database.DB.Create(&input).Error; err != nil {
		c.HTML(http.StatusInternalServerError, "list.html", gin.H{
			"Error": "failed to add sneaker",
		})
		return
	}

	c.Redirect(http.StatusOK, "/admin/products")
}

func UpdateSneaker(c *gin.Context) {
	id := c.Param("id")

	var sneaker models.Sneaker

	if err := database.DB.First(&sneaker, id).Error; err != nil {
		c.HTML(http.StatusNotFound, "list.html", gin.H{
			"Error": "sneaker not found",
		})
		return
	}

	if err := c.ShouldBindJSON(&sneaker); err != nil {
		c.HTML(http.StatusBadRequest, "list.html", gin.H{
			"Error": "invalid input",
		})
		return
	}

	if err := database.DB.Save(&sneaker).Error; err != nil {
		c.HTML(http.StatusInternalServerError, "list.html", gin.H{
			"Error": "failed to update sneaker",
		})
		return
	}

	c.Redirect(http.StatusOK, "/admin/products")

}

func DeleteSneaker(c *gin.Context) {
	id := c.Param("id")
	var sneaker models.Sneaker

	if err := database.DB.Delete(&sneaker, id).Error; err != nil {

		if err == gorm.ErrRecordNotFound {
			c.HTML(http.StatusInternalServerError,"list.html", gin.H{
				"Error": "sneaker not found",
			})
			return
		}
		c.HTML(http.StatusInternalServerError,"list.html", gin.H{
			"Error": "failed to delete sneaker",
		})
		return
	}

	c.Redirect(http.StatusOK, "/admin/products")
}
