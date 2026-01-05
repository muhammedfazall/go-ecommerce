package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/muhammedfazall/go-ecommerce/internal/database"
	"github.com/muhammedfazall/go-ecommerce/internal/models"
)

func GetAllSneakers(c *gin.Context) {
	var sneakers []models.Sneaker

	if err := database.DB.Order("created_at desc").Find(&sneakers).Error; err != nil {
		c.JSON(500,gin.H{
			"error":"failed to fetch sneakers",
		})
		return
	}
	c.JSON(http.StatusOK,sneakers)
}

func AddSneaker(c *gin.Context){
	var input models.Sneaker

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest,gin.H{
			"error":"invalid input",
		})
		return
	}

	if err:= database.DB.Create(&input).Error; err!= nil{
		c.JSON(http.StatusInternalServerError,gin.H{
			"error":"failed to add sneaker",
		})
		return
	}

	c.JSON(http.StatusOK, input)
}

func UpdateSneaker( c *gin.Context)  {
	id := c.Param("id")

	var sneaker models.Sneaker

	if err := database.DB.First(&sneaker, id).Error; err != nil{
		c.JSON(http.StatusNotFound,gin.H{"error":"sneaker not found"})
		return
	}

	if err := c.ShouldBindJSON(&sneaker); err != nil{
		c.JSON(http.StatusBadRequest,gin.H{"error":"invalid input"})
		return
	}

	if err := database.DB.Save(&sneaker).Error; err != nil {
		c.JSON(http.StatusInternalServerError,gin.H{
			"error":"failed to update sneaker",
		})
		return
	}

	c.JSON(http.StatusOK,sneaker)
}

func DeleteSneaker(c *gin.Context)  {
	id := c.Param("id")
	var sneaker models.Sneaker

	if err := database.DB.First(&sneaker, id).Error; err != nil {
		c.JSON(http.StatusInternalServerError,gin.H{
			"error":"sneaker not found",
		})
		return
	}

	sneaker.IsActive = false

	if err := database.DB.Save(&sneaker).Error; err != nil{
		c.JSON(http.StatusInternalServerError,gin.H{
			"error":"failed to delete sneaker",
		})
		return
	}

	c.JSON(http.StatusOK,gin.H{
		"message":"sneaker deactivated", //isActive -> false
	})
}