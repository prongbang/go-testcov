package model

// Aqi is model
type Aqi struct {
	Stations []struct {
		StationID   string `json:"stationID"`
		NameTH      string `json:"nameTH"`
		NameEN      string `json:"nameEN"`
		AreaTH      string `json:"areaTH"`
		AreaEN      string `json:"areaEN"`
		StationType string `json:"stationType"`
		Lat         string `json:"lat"`
		Long        string `json:"long"`
		AQILast     struct {
			Date string `json:"date"`
			Time string `json:"time"`
			PM25 struct {
				ColorID string `json:"color_id"`
				Aqi     string `json:"aqi"`
				Value   string `json:"value"`
			} `json:"PM25"`
			PM10 struct {
				ColorID string `json:"color_id"`
				Aqi     string `json:"aqi"`
				Value   string `json:"value"`
			} `json:"PM10"`
			O3 struct {
				ColorID string `json:"color_id"`
				Aqi     string `json:"aqi"`
				Value   string `json:"value"`
			} `json:"O3"`
			CO struct {
				ColorID string `json:"color_id"`
				Aqi     string `json:"aqi"`
				Value   string `json:"value"`
			} `json:"CO"`
			NO2 struct {
				ColorID string `json:"color_id"`
				Aqi     string `json:"aqi"`
				Value   string `json:"value"`
			} `json:"NO2"`
			SO2 struct {
				ColorID string `json:"color_id"`
				Aqi     string `json:"aqi"`
				Value   string `json:"value"`
			} `json:"SO2"`
			AQI struct {
				ColorID string `json:"color_id"`
				Aqi     string `json:"aqi"`
				Param   string `json:"param"`
			} `json:"AQI"`
		} `json:"AQILast"`
	} `json:"stations"`
}
