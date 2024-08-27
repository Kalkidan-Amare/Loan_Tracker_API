package domain

import (
	"time"
)

type Loan struct {
	ID         string    `json:"id" bson:"_id,omitempty"`
	UserID     string    `json:"user_id" bson:"user_id"`
	Amount     float64   `json:"amount" bson:"amount"`
	Status     string    `json:"status" bson:"status"`
	CreatedAt  time.Time `json:"created_at" bson:"created_at"`
	UpdatedAt  time.Time `json:"updated_at" bson:"updated_at"`
}

type LoanUsecaseInterface interface {
	CreateLoan(loan Loan) (string, error)
	GetLoanStatus(id string) (Loan, error)
	GetAllLoans(filter LoanFilter) ([]Loan, error)
	ApproveRejectLoan(id string, status string) error
	RemoveLoan(id string) error
}

type LoanRepositoryInteface interface {
	CreateLoan(loan domain.Loan) (string, error)
	GetLoanByID(id string) (domain.Loan, error)
	GetAllLoans(status, order string) ([]domain.Loan, error)
	UpdateLoanStatus(id, status string) error
	DeleteLoan(id string) error
}