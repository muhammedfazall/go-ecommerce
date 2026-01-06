package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/muhammedfazall/go-ecommerce/internal/database"
	"github.com/muhammedfazall/go-ecommerce/internal/models"
)

// all products
func GetAllSneakers(c *gin.Context) {
	var sneakers []models.Sneaker

	if err := database.DB.Order("created_at desc").Find(&sneakers).Error; err != nil {
		c.HTML(500, "products.html", gin.H{
			"Error": "failed to fetch sneakers",
		})
		return
	}
	c.HTML(http.StatusOK, "products.html", gin.H{
		"Sneakers": sneakers,
	})
}

//add product
func AddSneakerPage(c *gin.Context) {
	c.HTML(http.StatusOK, "product_add.html", nil)
}

func AddSneaker(c *gin.Context) {

	// Input struct
	var input struct {
		Name     string  `form:"name" binding:"required"`
		Brand    string  `form:"brand" binding:"required"`
		Gender   string  `form:"gender" binding:"required"`
		Price    float64 `form:"price" binding:"required"`
		Stock    int     `form:"stock" binding:"required"`
		ImageURL string  `form:"image_url"`
	}

	if err := c.ShouldBind(&input); err != nil {
		c.HTML(http.StatusBadRequest, "product_add.html", gin.H{
			"Error": "Invalid input",
		})
		return
	}

	// add explicitly
	sneaker := models.Sneaker{
		Name:     input.Name,
		Brand:    input.Brand,
		Gender:   input.Gender,
		Price:    input.Price,
		Stock:    input.Stock,
		ImageURL: input.ImageURL,
		IsActive: true,
	}

	if err := database.DB.Create(&sneaker).Error; err != nil {
		c.HTML(http.StatusInternalServerError, "product_add.html", gin.H{
			"Error": "Failed to add sneaker",
		})
		return
	}

	c.Redirect(http.StatusSeeOther, "/admin/products")
}

//view product

func ViewSneaker(c *gin.Context)  {
	id := c.Param("id")

	var sneaker models.Sneaker
	if err := database.DB.First(&sneaker, id).Error; err != nil {
		c.Redirect(http.StatusSeeOther, "/admin/products")
		return
	}

	c.HTML(http.StatusOK,"product_details.html",gin.H{
		"Sneaker":sneaker,
	})
}

// edit product
func EditSneaker(c *gin.Context) {
	id := c.Param("id")

	var sneaker models.Sneaker
	if err := database.DB.First(&sneaker, id).Error; err != nil {
		c.Redirect(http.StatusSeeOther, "/admin/products")
		return
	}

	c.HTML(http.StatusOK, "product_edit.html", gin.H{
		"Sneaker": sneaker,
	})
}

func UpdateSneaker(c *gin.Context) {
	id := c.Param("id")

	var sneaker models.Sneaker
	if err := database.DB.First(&sneaker, id).Error; err != nil {
		c.Redirect(http.StatusSeeOther, "/admin/products")
		return
	}

	// Input struct
	var input struct {
		Name     string  `form:"name" binding:"required"`
		Brand    string  `form:"brand" binding:"required"`
		Gender   string  `form:"gender"`
		Price    float64 `form:"price"`
		Stock    int     `form:"stock"`
		ImageURL string  `form:"image_url"`
		IsActive bool    `form:"is_active"`
	}

	if err := c.ShouldBind(&input); err != nil {
		c.HTML(http.StatusBadRequest, "product_edit.html", gin.H{
			"Error":   "Invalid input",
			"Sneaker": sneaker,
		})
		return
	}

	// Update fields explicitly
	sneaker.Name = input.Name
	sneaker.Brand = input.Brand
	sneaker.Gender = input.Gender
	sneaker.Price = input.Price
	sneaker.Stock = input.Stock
	sneaker.ImageURL = input.ImageURL
	sneaker.IsActive = input.IsActive

	if err := database.DB.Save(&sneaker).Error; err != nil {
		c.HTML(http.StatusInternalServerError, "product_edit.html", gin.H{
			"Error":   "Failed to update sneaker",
			"Sneaker": sneaker,
		})
		return
	}

	c.Redirect(http.StatusSeeOther, "/admin/products")
}

func DeleteSneaker(c *gin.Context) {
	id := c.Param("id")

	// Check existence
	var sneaker models.Sneaker
	if err := database.DB.Delete(&sneaker, id).Error; err != nil {
		c.Redirect(http.StatusSeeOther, "/admin/products")
		return
	}

	//delete
	if err := database.DB.Delete(&sneaker).Error; err != nil {
		c.Redirect(http.StatusSeeOther, "/admin/products")
		return
	}

	// redirect after delete
	c.Redirect(http.StatusOK, "/admin/products")
}
