package datasource

import (
	"github.com/prongbang/go-testcov/pm25/core"
	"github.com/prongbang/go-testcov/pm25/model"
)

// PmDataSource is the interface
type PmDataSource interface {
	GetAQI(url string) (model.Aqi, error)
}

type pmDataSource struct {
	Request core.HTTPRequest
}

// NewPmDataSource is new instance
func NewPmDataSource(req core.HTTPRequest) PmDataSource {
	return &pmDataSource{
		Request: req,
	}
}

func (ds *pmDataSource) GetAQI(url string) (model.Aqi, error) {

	response := new(model.Aqi)

	err := ds.Request.GetJSON(url, response)

	return *response, err
}
