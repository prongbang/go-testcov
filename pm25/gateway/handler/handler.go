package handler

import (
	"encoding/json"
	"net/http"

	"github.com/prongbang/go-testcov/pm25/domain"
)

// PmHandler is the interface
type PmHandler interface {
	GetAqi(w http.ResponseWriter, r *http.Request)
}

type pmHandler struct {
	UseCase domain.PmUseCase
}

// NewPmHandler is function new instance
func NewPmHandler(useCase domain.PmUseCase) PmHandler {
	return &pmHandler{
		UseCase: useCase,
	}
}

func (h *pmHandler) GetAqi(w http.ResponseWriter, r *http.Request) {

	url := "http://air4thai.pcd.go.th/forappV2/getAQI_JSON.php"

	data, _ := h.UseCase.GetAQI(url)

	res, _ := json.Marshal(data)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
