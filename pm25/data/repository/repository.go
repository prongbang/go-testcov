package repository

import (
	"github.com/prongbang/go-testcov/pm25/data/datasource"
	"github.com/prongbang/go-testcov/pm25/model"
)

// PmRepository is the interface
type PmRepository interface {
	GetAQI(url string) (model.Aqi, error)
}

type pmRepository struct {
	DataSource datasource.PmDataSource
}

// NewPmRepository is function new instance
func NewPmRepository(dataSource datasource.PmDataSource) PmRepository {
	return &pmRepository{
		DataSource: dataSource,
	}
}

func (repo *pmRepository) GetAQI(url string) (model.Aqi, error) {

	return repo.DataSource.GetAQI(url)
}
