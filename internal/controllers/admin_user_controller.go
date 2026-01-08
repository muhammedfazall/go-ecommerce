package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/muhammedfazall/go-ecommerce/internal/database"
	"github.com/muhammedfazall/go-ecommerce/internal/models"
)

//list users
func GetUsers(c *gin.Context) {
	var users []models.User

	search := c.Query("search")
	status := c.Query("status")
	blocked := c.Query("blocked")

	query := database.DB.Where("role = ?", "user")

	if search != "" {
		query = query.Where(
			"username ILIKE ? OR email ILIKE ? OR CAST(id AS TEXT) ILIKE ?",
			"%"+search+"%",
			"%"+search+"%",
			"%"+search+"%",
		)
	}

	if status != "" {
		query = query.Where("status = ?", status)
	}

	if blocked != "" {
		query = query.Where("is_blocked = ?", blocked)
	}

	query.Find(&users)

	c.HTML(200, "users.html", gin.H{
		"Users":     users,
		"Search": search,
		"Searched":  search != "",
	})
}


//block/unblock
func ToggleBlockUser(c *gin.Context) {
	id := c.Param("id")

	var user models.User
	if err := database.DB.First(&user, id).Error; err != nil {
		c.Redirect(302, "/admin/users")
		return
	}

	user.IsBlocked = !user.IsBlocked
	database.DB.Save(&user)

	c.Redirect(302, "/admin/users")
}

// change role
func ChangeUserRole(c *gin.Context) {
	id := c.Param("id")
	newRole := c.PostForm("role")

	if newRole != "user" && newRole != "admin" {
		c.Redirect(302, "/admin/users")
		return
	}

	database.DB.Model(&models.User{}).
		Where("id = ?", id).
		Update("role", newRole)

	// Auto-move between sections
	if newRole == "admin" {
		c.Redirect(302, "/admin/admins")
		return
	}

	c.Redirect(302, "/admin/users")
}
