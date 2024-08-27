package usecases

// import (
// 	"loan-tracker/domain"

// 	"go.mongodb.org/mongo-driver/bson/primitive"
// )

// type AdminUsecase struct {
// 	userRepo  domain.UserRepositoryInterface
// }

// func NewAdminUsecase(ur domain.UserRepositoryInterface) *UserUsecase {
// 	return &UserUsecase{
// 		userRepo:  ur,
// 	}
// }


// func (u *AdminUsecase) GetAllUsers() ([]*domain.User, error) {
// 	users, err := u.userRepo.GetAllUsers()
// 	return users, err
// }

// // DeleteUser deletes a user by ID
// func (u *AdminUsecase) DeleteUser(objectID primitive.ObjectID) error {
// 	return u.userRepo.DeleteUser(objectID)
// }

import (
	"loan-tracker/domain"
	"loan-tracker/repositories"
)

type AdminUsecase interface {
	ViewAllLoans(status, order string) ([]domain.Loan, error)
	UpdateLoanStatus(id, status string) error
	DeleteLoan(id string) error
	ViewSystemLogs() ([]domain.Log, error)
}

type adminUsecase struct {
	loanRepo repositories.LoanRepository
	logRepo  repositories.LogRepository
}

func NewAdminUsecase(loanRepo repositories.LoanRepository, logRepo repositories.LogRepository) AdminUsecase {
	return &adminUsecase{loanRepo: loanRepo, logRepo: logRepo}
}

func (au *adminUsecase) ViewAllLoans(status, order string) ([]domain.Loan, error) {
	return au.loanRepo.GetAllLoans(status, order)
}

func (au *adminUsecase) UpdateLoanStatus(id, status string) error {
	return au.loanRepo.UpdateLoanStatus(id, status)
}

func (au *adminUsecase) DeleteLoan(id string) error {
	return au.loanRepo.DeleteLoan(id)
}

func (au *adminUsecase) ViewSystemLogs() ([]domain.Log, error) {
	return au.logRepo.GetAllLogs()
}
