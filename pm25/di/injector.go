package di

import (
	"net/http"
	"time"

	"github.com/prongbang/go-testcov/pm25/core"
	"github.com/prongbang/go-testcov/pm25/data/datasource"
	"github.com/prongbang/go-testcov/pm25/data/repository"
	"github.com/prongbang/go-testcov/pm25/domain"
	"github.com/prongbang/go-testcov/pm25/gateway/handler"
	"github.com/prongbang/go-testcov/pm25/gateway/route"
)

// ProvideHTTPRequest is function provide instance
func ProvideHTTPRequest() core.HTTPRequest {

	return core.NewHTTPRequest(&http.Client{Timeout: 60 * time.Second})
}

// ProvidePmDataSource is function provide instance
func ProvidePmDataSource() datasource.PmDataSource {
	return datasource.NewPmDataSource(ProvideHTTPRequest())
}

// ProvidePmRepository is function provide instance
func ProvidePmRepository() repository.PmRepository {
	return repository.NewPmRepository(ProvidePmDataSource())
}

// ProvidePmUseCase is function provide instance
func ProvidePmUseCase() domain.PmUseCase {
	return domain.NewPmUseCase(ProvidePmRepository())
}

// ProvidePmHandler is function provide instance
func ProvidePmHandler() handler.PmHandler {
	return handler.NewPmHandler(ProvidePmUseCase())
}

// ProvidePmRoute is function provide instance
func ProvidePmRoute() route.PmRoute {
	return route.NewPmRoute(ProvidePmHandler())
}
