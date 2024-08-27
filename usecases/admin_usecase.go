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