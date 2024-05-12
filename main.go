package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/fatih/color"
	"github.com/joho/godotenv"
)

type Weather struct {
	Location struct {
		Name      string `json:"name"`
		Region    string `json:"Region"`
		Country   string `json:"country"`
		LocalTime string `json:"localtime"`
	} `json:"location"`
	Current struct {
		TempC     float64 `json:"temp_c"`
		Condition struct {
			Text string `json:"text"`
		} `json:"condition"`
	} `json:"current"`
	Forecast struct {
		ForecastDay []struct {
			Day struct {
				MaxTempC  float64 `json:"maxtemp_c"`
				MinTempC  float64 `json:"mintemp_c"`
				Condition struct {
					Text string `json:"text"`
				} `json:"condition"`
				UV float64 `json:"uv"`
			} `json:"day"`
			Hour []struct {
				TimeEpoch int64   `json:"time_epoch"`
				TempC     float64 `json:"temp_c"`
				Condition struct {
					Text string `json:"text"`
				} `json:"condition"`
				ChanceOfRain float64 `json:"chance_of_rain"`
			} `json:"hour"`
		} `json:"forecastday"`
	} `json:"forecast"`
}

func main() {
	q := "Bengaluru"

	if len(os.Args) >= 2 {
		q = os.Args[1]
	}
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error while reading env file: %v", err)

	}
	env := os.Getenv("ApiKey")
	endpoint := "http://api.weatherapi.com/v1/forecast.json?key=" + env + "&q=" + q + "&days=1&aqi=yes&alerts=no"

	res, err := http.Get(endpoint)

	if err != nil {
		//panic(err)
		log.Fatalf("Error while fetching the endpoint: %v", err)
	}

	defer res.Body.Close()

	if res.StatusCode != 200 {
		panic("Weather API not found!")
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		//panic(err)
		log.Fatalf("Unable to Read the data from json: %v", err)
	}
	//fmt.Println(string(body))
	var weather Weather
	err = json.Unmarshal(body, &weather)
	if err != nil {
		log.Fatalf("Unable to Unmarshal the json: %v", err)
	}
	// fmt.Println(weather)

	location, current, hours := weather.Location, weather.Current, weather.Forecast.ForecastDay[0].Hour

	fmt.Printf(
		"%s, %s: %.0f°C, %s\nLocal Time: %s",
		location.Name, location.Country, current.TempC, current.Condition.Text, location.LocalTime,
	)
	fmt.Println("\nToday's upcoming forcast: ")
	for _, hour := range hours {

		date := time.Unix(hour.TimeEpoch, 0)
		//Comment the below 3 lines of code if you want to see the full day forcast
		if date.Before(time.Now()) {
			continue
		}

		message := fmt.Sprintf(
			"%s - %.0f°C, %.0f%%, %s\n",
			date.Format("15.04"),
			hour.TempC,
			hour.ChanceOfRain,
			hour.Condition.Text,
		)
		if hour.ChanceOfRain < 75 {
			fmt.Print(message)
		} else {
			color.Red(message)
		}

	}
	fmt.Println()
}
