package core

import (
	"encoding/json"
	"net/http"
)

// HTTPRequest is the interface
type HTTPRequest interface {
	GetJSON(url string, target interface{}) error
}

type httpRequest struct {
	Client *http.Client
}

// NewHTTPRequest is new instance
func NewHTTPRequest(client *http.Client) HTTPRequest {
	return &httpRequest{
		Client: client,
	}
}

func (req *httpRequest) GetJSON(url string, target interface{}) (err error) {
	r, _ := req.Client.Get(url)

	return json.NewDecoder(r.Body).Decode(target)
}
