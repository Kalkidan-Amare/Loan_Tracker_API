package repositories

import (
	"context"
	"loan-tracker/domain"
	"go.mongodb.org/mongo-driver/mongo"
)

type LogRepository interface {
	LogEvent(log domain.Log) error
	GetAllLogs() ([]domain.Log, error)
}

type logRepository struct {
	collection *mongo.Collection
}

func NewLogRepository(collection *mongo.Collection) LogRepository {
	return &logRepository{collection: collection}
}

func (lr *logRepository) LogEvent(log domain.Log) error {
	_, err := lr.collection.InsertOne(context.Background(), log)
	return err
}

func (lr *logRepository) GetAllLogs() ([]domain.Log, error) {
	cursor, err := lr.collection.Find(context.Background(), bson.M{})
	if err != nil {
		return nil, err
	}

	var logs []domain.Log
	err = cursor.All(context.Background(), &logs)
	return logs, err
}
