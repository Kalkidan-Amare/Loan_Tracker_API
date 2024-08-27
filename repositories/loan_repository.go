package repositories

import (
	"context"
	"loan-tracker/domain"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type LoanRepository struct {
	collection *mongo.Collection
}

func NewLoanRepository(collection *mongo.Collection) *LoanRepository {
	return &LoanRepository{collection: collection}
}

func (lr *LoanRepository) CreateLoan(loan domain.Loan) (string, error) {
	_, err := lr.collection.InsertOne(context.Background(), loan)
	if err != nil {
		return "", err
	}
	return "Loan application submitted successfully", nil
}

func (lr *LoanRepository) GetLoanByID(id string) (domain.Loan, error) {
	var loan domain.Loan
	err := lr.collection.FindOne(context.Background(), bson.M{"_id": id}).Decode(&loan)
	return loan, err
}

func (lr *LoanRepository) GetAllLoans(status, order string) ([]domain.Loan, error) {
	filter := bson.M{}
	if status != "all" {
		filter["status"] = status
	}

	opts := options.Find()
	if order == "desc" {
		opts.SetSort(bson.D{{Key: "created_at", Value: -1}})
	} else {
		opts.SetSort(bson.D{{Key: "created_at", Value: 1}})
	}

	cursor, err := lr.collection.Find(context.Background(), filter, opts)
	if err != nil {
		return nil, err
	}

	var loans []domain.Loan
	err = cursor.All(context.Background(), &loans)
	return loans, err
}

func (lr *LoanRepository) UpdateLoanStatus(id, status string) error {
	_, err := lr.collection.UpdateOne(context.Background(), bson.M{"_id": id}, bson.M{"$set": bson.M{"status": status}})
	return err
}

func (lr *LoanRepository) DeleteLoan(id string) error {
	_, err := lr.collection.DeleteOne(context.Background(), bson.M{"_id": id})
	return err
}
