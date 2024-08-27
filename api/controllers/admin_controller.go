package controllers

import (
	"net/http"

	"loan-tracker/domain"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type AdminController struct {
	uUsecase domain.UserUsecaseInterface
	adminUsecase domain.LoanUsecaseInterface
}

func NewAdminController(uu domain.UserUsecaseInterface, au domain.LoanUsecaseInterface) *AdminController {
	return &AdminController{
		uUsecase: uu,
		adminUsecase: au,
	}
}

// GetAllUsers handles the GET request to retrieve all users
func (ac *AdminController) GetAllUsers(c *gin.Context) {
	users, err := ac.uUsecase.GetAllUsers()
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


func (ac *AdminController) ViewAllLoans(c *gin.Context) {
	status := c.DefaultQuery("status", "all")
	order := c.DefaultQuery("order", "asc")

	loans, err := ac.adminUsecase.ViewAllLoans(status, order)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, loans)
}

func (ac *AdminController) UpdateLoanStatus(c *gin.Context) {
	id := c.Param("id")
	var input struct {
		Status string `json:"status"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := ac.adminUsecase.UpdateLoanStatus(id, input.Status)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Loan status updated"})
}

func (ac *AdminController) DeleteLoan(c *gin.Context) {
	id := c.Param("id")

	err := ac.adminUsecase.DeleteLoan(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Loan deleted successfully"})
}

func (ac *AdminController) ViewSystemLogs(c *gin.Context) {
	logs, err := ac.adminUsecase.ViewSystemLogs()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, logs)
}
