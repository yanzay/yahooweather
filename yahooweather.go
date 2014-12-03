package yahooweather

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
)

// Current weather condition
type ConditionItem struct {
	Code int `json:",string"`
	Temp int `json:",string"`
}

// Forecast for one date.
// * Code - weather code
// * High - high bound of temperature
// * Low - low bound of temperature
// * Date - string representation of date
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
					Condition ConditionItem
					Forecast  []ForecastItem
				}
			}
		}
	}
}

func GetWeather(id int64) (ConditionItem, []ForecastItem, error) {
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

func makeRequest(request string) (ConditionItem, []ForecastItem, error) {
	log.Printf("Requesting %s", request)
	resp, err := http.Get(request)
	if err != nil {
		return ConditionItem{}, nil, err
	}
	defer resp.Body.Close()
	r := new(response)
	json.NewDecoder(resp.Body).Decode(r)
	item := r.Query.Results.Channel.Item
	return item.Condition, item.Forecast, nil
}
