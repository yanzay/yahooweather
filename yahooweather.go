package yahooweather

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
)

// Forecast for one date.
// Code - weather code
// High - high bound of temperature
// Low - low bound of temperature
// Date - string representation of date
type ForecastItem struct {
	Code int `json:",string"`
	High int `json:",string"`
	Low  int `json:",string"`
	Date string
}

type response struct {
	Query struct {
		Results struct {
			Channel struct {
				Item struct {
					Forecast []ForecastItem
				}
			}
		}
	}
}

// Get forecasts from Yahoo. Returns slice of ForecastItem's
func GetForecasts(id int64) ([]ForecastItem, error) {
	request := prepareRequest(id)
	return makeRequest(request)
}

func prepareRequest(id int64) (request string) {
	request = "https://query.yahooapis.com/v1/public/yql?q="
	query := fmt.Sprintf("select * from weather.forecast where woeid=%d and u='c'", id)
	request += url.QueryEscape(query)
	request += "&format=json"
	return
}

func makeRequest(request string) ([]ForecastItem, error) {
	resp, err := http.Get(request)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	r := new(response)
	json.NewDecoder(resp.Body).Decode(r)
	return r.Query.Results.Channel.Item.Forecast, nil
}
