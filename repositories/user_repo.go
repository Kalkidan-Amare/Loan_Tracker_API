package repositories

import (
	"context"
	"loan-tracker/domain"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	// "gorm.io/gorm"
)

// type UserRepository interface {
//     Create(user *domain.User) error
//     GetByEmail(email string) (*domain.User, error)
//     Update(user *domain.User) error
//     GetAll() ([]*domain.User, error)
//     DeleteByID(id string) error
// }

type UserRepository struct {
	collection domain.Collection
}

func NewUserRepository(col domain.Collection) *UserRepository {
	return &UserRepository{
		collection: col,
	}
}

func (r *UserRepository) CreateUser(user *domain.User) error {
    ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	_, err := r.collection.InsertOne(ctx, user)
	return err
}

func (r *UserRepository) GetUserByEmail(email string) (*domain.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	var user domain.User
	err := r.collection.FindOne(ctx, bson.M{"email": email}).Decode(&user)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *UserRepository) GetUserByUsername(username string) (*domain.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	var user domain.User
	err := r.collection.FindOne(ctx, bson.M{"name": username}).Decode(&user)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *UserRepository) GetUserByID(id primitive.ObjectID) (*domain.User, error) {
	var user domain.User
	err := r.collection.FindOne(context.TODO(), bson.M{"_id": id}).Decode(&user)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *UserRepository) UpdateUser(username string, user *domain.User) error {
	result := r.collection.FindOneAndUpdate(context.TODO(), bson.M{"name": username}, bson.M{"$set": user})
	
	return result.Err()
}

func (r *UserRepository) GetAllUsers() ([]*domain.User, error) {
	var users []*domain.User
	cursor, err := r.collection.Find(context.TODO(), bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.TODO())

	err = cursor.All(context.TODO(), &users)

	return users, err
}

func (r *UserRepository) DeleteUser(id primitive.ObjectID) error {
	_, err := r.collection.DeleteOne(context.TODO(), bson.M{"_id": id})
	return err
}
