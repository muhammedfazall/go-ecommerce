package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/muhammedfazall/go-ecommerce/services"
)

type RegisterRequest struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func Register(c *gin.Context) {
	var req RegisterRequest


	if err := c.ShouldBindJSON(&req) ; err != nil {
		c.JSON(http.StatusBadRequest,gin.H{
			"error":"invalid request body",
		})
		return
	}

	if req.Name == "" || req.Email == "" || req.Password == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "name, email and password are required",
		})
		return
	}

	err := services.RegisterUser(req.Name,req.Email,req.Password)
	if err != nil{
		c.JSON(http.StatusBadRequest,gin.H{
			"error":err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated,gin.H{
		"message":"user registered successfully",
	})
}
