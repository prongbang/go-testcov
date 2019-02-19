package route

import (
	"net/http"

	"github.com/prongbang/go-testcov/pm25/gateway/handler"
)

// PmRoute is the interface
type PmRoute interface {
	Initial(h *http.ServeMux)
}

type pmRoute struct {
	Handle handler.PmHandler
}

// NewPmRoute is function new instance
func NewPmRoute(handle handler.PmHandler) PmRoute {
	return &pmRoute{
		Handle: handle,
	}
}

func (r *pmRoute) Initial(h *http.ServeMux) {
	h.HandleFunc("/aqi", r.Handle.GetAqi)
}
