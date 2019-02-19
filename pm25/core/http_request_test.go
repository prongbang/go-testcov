package core_test

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/prongbang/go-testcov/pm25/core"
	"github.com/prongbang/go-testcov/pm25/model"
)

func responseMock() string {
	return `{ "stations": [{ "stationID": "03t", "nameTH": "ริมถนนทางหลวงหมายเลข 3902 ", "nameEN": "Highway NO.3902 km.13 +600", "areaTH": "ริมถนนกาญจนาภิเษก เขตบางขุนเทียน, กรุงเทพฯ", "areaEN": "Kanchanaphisek Rd, Bang Khun Thian, Bangkok", "stationType": "GROUND", "lat": "13.636514", "long": "100.414262", "AQILast": { "date": "2019-02-17", "time": "14:00", "PM25": { "color_id": "2", "aqi": "26", "value": "26" }, "PM10": { "color_id": "2", "aqi": "26", "value": "51" }, "O3": { "color_id": "1", "aqi": "8", "value": "11" }, "CO": { "color_id": "1", "aqi": "3", "value": "0.54" }, "NO2": { "color_id": "1", "aqi": "13", "value": "30" }, "SO2": { "color_id": "0", "aqi": "-999", "value": "-1" }, "AQI": { "color_id": "2", "aqi": "26", "param": "PM10" } } }] }`
}

// RoundTripFunc .
type RoundTripFunc func(req *http.Request) *http.Response

// RoundTrip .
func (f RoundTripFunc) RoundTrip(req *http.Request) (*http.Response, error) {
	return f(req), nil
}

// NewTestClient returns *http.Client with Transport replaced to avoid making real calls
func NewTestClient(fn RoundTripFunc) *http.Client {
	return &http.Client{
		Transport: RoundTripFunc(fn),
	}
}

func TestGetJSONSuccess(t *testing.T) {

	client := NewTestClient(func(req *http.Request) *http.Response {

		if req.URL.String() != "http://air4thai.pcd.go.th/forappV2/getAQI_JSON.php" {
			t.Error("Not found")
		}

		return &http.Response{
			StatusCode: http.StatusOK,
			Body:       ioutil.NopCloser(bytes.NewBufferString(responseMock())),
			Header:     make(http.Header),
		}
	})

	request := core.NewHTTPRequest(client)
	url := "http://air4thai.pcd.go.th/forappV2/getAQI_JSON.php"
	data := new(model.Aqi)
	request.GetJSON(url, data)

	if len(data.Stations) == 0 {
		t.Error("Size is 0")
	}

	if data.Stations[0].StationID != "03t" {
		t.Error("StationID is not 03t")
	}
}

func TestGetJSONFailure(t *testing.T) {

	// Start a local HTTP server
	server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		// Send response to be tested
		rw.Write([]byte(`OK`))
	}))
	// Close the server when test finishes
	defer server.Close()

	request := core.NewHTTPRequest(server.Client())
	url := "http://test.com/#q=1"
	data := new(model.Aqi)
	err := request.GetJSON(url, data)

	if err == nil {
		t.Error("Size is not 0", err)
	}

}
