package domain

import (
	"github.com/prongbang/go-testcov/pm25/data/repository"
	"github.com/prongbang/go-testcov/pm25/model"
)

// PmUseCase is the interface
type PmUseCase interface {
	GetAQI(url string) (model.Aqi, error)
}

type pmUseCase struct {
	Repo repository.PmRepository
}

// NewPmUseCase is function new instance
func NewPmUseCase(repo repository.PmRepository) PmUseCase {
	return &pmUseCase{
		Repo: repo,
	}
}

func (useCase *pmUseCase) GetAQI(url string) (model.Aqi, error) {

	return useCase.Repo.GetAQI(url)
}
