package usecases

import (
	"loan-tracker/domain"
	"loan-tracker/repositories"
)

type LoanUsecase struct {
	loanRepo domain.LoanRepositoryInteface
}

func NewLoanUsecase(loanRepo domain.LoanRepositoryInteface) *LoanUsecase {
	return &LoanUsecase{
		loanRepo: loanRepo,
	}
}

func (u *LoanUsecase) CreateLoan(loan domain.Loan) (string, error) {
	return u.loanRepo.CreateLoan(loan)
}

func (u *LoanUsecase) GetLoanStatus(id string) (domain.Loan, error) {
	return u.loanRepo.GetLoanByID(id)
}

func (u *LoanUsecase) GetAllLoans(filter domain.LoanFilter) ([]domain.Loan, error) {
	return u.loanRepo.GetAllLoans(filter)
}

func (u *LoanUsecase) ApproveRejectLoan(id string, status string) error {
	return u.loanRepo.UpdateLoanStatus(id, status)
}

func (u *LoanUsecase) RemoveLoan(id string) error {
	return u.loanRepo.DeleteLoan(id)
}
