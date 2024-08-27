package domain

import "time"

type Log struct {
	ID        string    `json:"id" bson:"_id,omitempty"`
	Event     string    `json:"event" bson:"event"`
	UserID    string    `json:"user_id,omitempty" bson:"user_id,omitempty"`
	Timestamp time.Time `json:"timestamp" bson:"timestamp"`
}

type LogUsecaseInterface interface {
	CreateLog(log Log) error
	FetchAllLogs() ([]Log, error)
}


type LogRepositoryInterface interface {
	SaveLog(log Log) error
	GetAllLogs() ([]Log, error)
}