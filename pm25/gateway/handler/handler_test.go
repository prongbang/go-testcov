package handler_test

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/prongbang/go-testcov/pm25/domain"
	"github.com/prongbang/go-testcov/pm25/gateway/handler"
	"github.com/prongbang/go-testcov/pm25/model"
)

var responseJson = `{ "stations": [{ "stationID": "03t", "nameTH": "ริมถนนทางหลวงหมายเลข 3902 ", "nameEN": "Highway NO.3902 km.13 +600", "areaTH": "ริมถนนกาญจนาภิเษก เขตบางขุนเทียน, กรุงเทพฯ", "areaEN": "Kanchanaphisek Rd, Bang Khun Thian, Bangkok", "stationType": "GROUND", "lat": "13.636514", "long": "100.414262", "AQILast": { "date": "2019-02-17", "time": "14:00", "PM25": { "color_id": "2", "aqi": "26", "value": "26" }, "PM10": { "color_id": "2", "aqi": "26", "value": "51" }, "O3": { "color_id": "1", "aqi": "8", "value": "11" }, "CO": { "color_id": "1", "aqi": "3", "value": "0.54" }, "NO2": { "color_id": "1", "aqi": "13", "value": "30" }, "SO2": { "color_id": "0", "aqi": "-999", "value": "-1" }, "AQI": { "color_id": "2", "aqi": "26", "param": "PM10" } } }] }`

func responseMock() string {
	return responseJson
}

type pmUseCaseMock struct{}

func NewPmUseCaseMock() domain.PmUseCase {
	return &pmUseCaseMock{}
}

func (uc *pmUseCaseMock) GetAQI(url string) (model.Aqi, error) {

	var data model.Aqi
	if responseMock() == "" {

		return data, errors.New("404 Not Found")
	}

	err := json.Unmarshal([]byte(responseMock()), &data)

	return data, err
}

func TestGetAQI(t *testing.T) {

	h := http.NewServeMux()

	useCaseMock := NewPmUseCaseMock()
	handle := handler.NewPmHandler(useCaseMock)

	// test server
	ts := httptest.NewServer(h)
	defer ts.Close()

	t.Run("get-aqi-success", func(t *testing.T) {
		req, _ := http.NewRequest("GET", "", nil)
		w := httptest.NewRecorder()
		handle.GetAqi(w, req)

		aqi := model.Aqi{}
		json.Unmarshal([]byte(w.Body.String()), &aqi)
		if aqi.Stations == nil {
			t.Errorf("Got %s, Want %s", aqi.Stations, responseMock())
		}
	})

	t.Run("get-aqi-notfound", func(t *testing.T) {
		req, _ := http.NewRequest("GET", "", nil)
		w := httptest.NewRecorder()

		responseJson = ""
		handle.GetAqi(w, req)

		var aqi model.Aqi
		json.Unmarshal([]byte(w.Body.String()), &aqi)
		if len(aqi.Stations) != 0 {
			t.Errorf("Got %s, Want %s", aqi.Stations, responseMock())
		}
	})
}
