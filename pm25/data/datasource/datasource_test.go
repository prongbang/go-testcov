package datasource_test

import (
	"encoding/json"
	"testing"

	"github.com/prongbang/go-testcov/pm25/core"
	"github.com/prongbang/go-testcov/pm25/data/datasource"
)

func responseMock() string {
	return `{ "stations": [{ "stationID": "03t", "nameTH": "ริมถนนทางหลวงหมายเลข 3902 ", "nameEN": "Highway NO.3902 km.13 +600", "areaTH": "ริมถนนกาญจนาภิเษก เขตบางขุนเทียน, กรุงเทพฯ", "areaEN": "Kanchanaphisek Rd, Bang Khun Thian, Bangkok", "stationType": "GROUND", "lat": "13.636514", "long": "100.414262", "AQILast": { "date": "2019-02-17", "time": "14:00", "PM25": { "color_id": "2", "aqi": "26", "value": "26" }, "PM10": { "color_id": "2", "aqi": "26", "value": "51" }, "O3": { "color_id": "1", "aqi": "8", "value": "11" }, "CO": { "color_id": "1", "aqi": "3", "value": "0.54" }, "NO2": { "color_id": "1", "aqi": "13", "value": "30" }, "SO2": { "color_id": "0", "aqi": "-999", "value": "-1" }, "AQI": { "color_id": "2", "aqi": "26", "param": "PM10" } } }] }`
}

type httpRequestMock struct {
}

func NewHTTPRequestMock() core.HTTPRequest {
	return &httpRequestMock{}
}

func (h *httpRequestMock) GetJSON(url string, target interface{}) error {

	return json.Unmarshal([]byte(responseMock()), target)
}

func TestGetAQI(t *testing.T) {

	req := NewHTTPRequestMock()
	dataSource := datasource.NewPmDataSource(req)
	response, _ := dataSource.GetAQI("http://air4thai.pcd.go.th/forappV2/getAQI_JSON.php")

	if len(response.Stations) == 0 {
		t.Error("Is Empty")
	}
}
