package controllers

import (
	"net/http"
	"loan-tracker/domain"
	"github.com/gin-gonic/gin"
)

type LoanController struct {
	loanUsecase domain.LoanUsecaseInterface
}

func NewLoanController(lUsecase domain.LoanUsecaseInterface) *LoanController {
	return &LoanController{loanUsecase: lUsecase}
}

func (lc *LoanController) ApplyForLoan(c *gin.Context) {
	var loan domain.Loan
	if err := c.ShouldBindJSON(&loan); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	status, err := lc.loanUsecase.ApplyForLoan(loan)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": status})
}

func (lc *LoanController) GetLoanStatus(c *gin.Context) {
	id := c.Param("id")

	loan, err := lc.loanUsecase.GetLoanStatus(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, loan)
}


func (lc *LoanController) ViewAllLoans(c *gin.Context) {
	status := c.DefaultQuery("status", "all")
	order := c.DefaultQuery("order", "asc")

	loans, err := lc.loanUsecase.ViewAllLoans(status, order)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, loans)
}

func (lc *LoanController) UpdateLoanStatus(c *gin.Context) {
	id := c.Param("id")
	var input struct {
		Status string `json:"status"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := lc.loanUsecase.UpdateLoanStatus(id, input.Status)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Loan status updated"})
}

func (lc *LoanController) DeleteLoan(c *gin.Context) {
	id := c.Param("id")

	err := lc.loanUsecase.DeleteLoan(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Loan deleted successfully"})
}

func (lc *LoanController) ViewSystemLogs(c *gin.Context) {
	logs, err := lc.loanUsecase.ViewSystemLogs()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, logs)
}