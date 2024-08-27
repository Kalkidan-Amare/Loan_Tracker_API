package usecases

import (
	"loan-tracker/domain"
)

type LogUsecase struct {
	logRepo domain.LogRepositoryInterface
}

func NewLogUsecase(logRepo domain.LogRepositoryInterface) *LogUsecase {
	return &LogUsecase{
		logRepo: logRepo,
	}
}

func (u *LogUsecase) CreateLog(log domain.Log) error {
	return u.logRepo.SaveLog(log)
}

func (u *LogUsecase) FetchAllLogs() ([]domain.Log, error) {
	return u.logRepo.GetAllLogs()
}
