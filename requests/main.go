package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gin-gonic/gin"
)



//external data
type HourlyData struct {
	Time           []string  `json:"time"`
	Temperature2M  []float64 `json:"temperature_2m"`
}

type HourlyUnits struct {
	Temperature2M string `json:"temperature_2m"`
}

type CurrentWeather struct {
	Time          string  `json:"time"`
	Temperature   float64 `json:"temperature"`
	WeatherCode   int     `json:"weathercode"`
	WindSpeed     float64 `json:"windspeed"`
	WindDirection int     `json:"winddirection"`
}

type WeatherData struct {
	Latitude            float64       `json:"latitude"`
	Longitude           float64       `json:"longitude"`
	Elevation           float64       `json:"elevation"`
	GenerationTimeMs    float64       `json:"generationtime_ms"`
	UTCOffsetSeconds    int           `json:"utc_offset_seconds"`
	Timezone            string        `json:"timezone"`
	TimezoneAbbreviation string        `json:"timezone_abbreviation"`
	Hourly              HourlyData    `json:"hourly"`
	HourlyUnits         HourlyUnits   `json:"hourly_units"`
	CurrentWeather      CurrentWeather `json:"current_weather"`
}


//routes
func main(){
	router := gin.Default()
	router.GET("https://api.open-meteo.com/v1/forecast?latitude=52.52&longitude=13.41&hourly=temperature_2m")


var jsonString = `{
	
	"latitude": 52.52,
	"longitude": 13.419,
	"elevation": 44.812,
	"generationtime_ms": 2.2119,
	"utc_offset_seconds": 0,
	"timezone": "Europe/London",
	"timezone_abbreviation": "GMT",
	"hourly": {
		"time": [
			"2023-09-13T00:00",
			"2023-09-13T01:00",
			"2023-09-13T02:00",
			"2023-09-13T03:00",
			"2023-09-13T04:00",
			"2023-09-13T05:00",
			"2023-09-13T06:00",
			"2023-09-13T07:00",
			"2023-09-13T08:00",
			"2023-09-13T09:00",
			"2023-09-13T10:00",
		],
		"temperature_2m": [
			18.8,
			18.6,
			18.4,
			18.2,
			17.7,
			17.1,
			17.7,
			18.9,
			19.6,
			20.1,
			20.8,
		],
	}

	"hourly_units": {
		"temperature_2m": "°C"
	  },

	"current_weather": {
		"time": "2022-07-01T09:00",
		"temperature": 13.3,
		"weathercode": 3,
		"windspeed": 10.3,
		"winddirection": 262
	  }
}`

	
var weatherData WeatherData
	err := json.Unmarshal([]byte(jsonString), &weatherData)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}


jsonString = `{
		"latitude": 52.52,
		"longitude": 13.419,
		"elevation": 44.812,
		"generationtime_ms": 2.2119,
		"utc_offset_seconds": 0,
		"timezone": "Europe/London",
		"timezone_abbreviation": "GMT",
		"hourly": {
			"time": [
				"2023-09-13T00:00",
				"2023-09-13T01:00",
				"2023-09-13T02:00",
				"2023-09-13T03:00",
				"2023-09-13T04:00",
				"2023-09-13T05:00",
				"2023-09-13T06:00",
				"2023-09-13T07:00",
				"2023-09-13T08:00",
				"2023-09-13T09:00",
				"2023-09-13T10:00",
			],
			"temperature_2m": [
				18.8,
				18.6,
				18.4,
				18.2,
				17.7,
				17.1,
				17.7,
				18.9,
				19.6,
				20.1,
				20.8,
			],
		}

		"hourly_units": {
			"temperature_2m": "°C"
		  },

		"current_weather": {
			"time": "2022-07-01T09:00",
			"temperature": 13.3,
			"weathercode": 3,
			"windspeed": 10.3,
			"winddirection": 262
		  }

	}`

jsonValue, _ := json.Marshal(jsonString)


response, err := http.Post("https://api.open-meteo.com/v1/forecast?latitude=52.52&longitude=13.41&hourly=temperature_2m",
	 "application/json", bytes.NewBuffer(jsonValue))
	 if err != nil {
        fmt.Printf("The HTTP request failed with error %s\n", err)
    } else {
        data, _ := ioutil.ReadAll(response.Body)
        fmt.Println(string(data))
    }

response, err = http.Get("https://api.open-meteo.com/v1/forecast?latitude=52.52&longitude=13.41&hourly=temperature_2m")
	if err != nil {
		fmt.Printf("The HTTP request faild error %s\n", err)
	} else{
		data, _ :=ioutil.ReadAll(response.Body)

		fmt.Println(string (data))
	}
}