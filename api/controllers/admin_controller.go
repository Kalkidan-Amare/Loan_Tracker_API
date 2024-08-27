package controllers

import (
	"net/http"

	"loan-tracker/domain"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type AdminController struct {
	adminUsecase domain.UserUsecaseInterface
}

func NewAdminController(au domain.UserUsecaseInterface) *AdminController {
	return &AdminController{
		adminUsecase: au,
	}
}

// GetAllUsers handles the GET request to retrieve all users
func (ac *AdminController) GetAllUsers(c *gin.Context) {
	users, err := ac.adminUsecase.GetAllUsers()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, users)
}

// DeleteUser handles the DELETE request to delete a user by ID
func (ac *AdminController) DeleteUser(c *gin.Context) {
	id, err := primitive.ObjectIDFromHex(c.Param("id")) // should be updated
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}
	err = ac.adminUsecase.DeleteUser(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "User deleted successfully"})
}