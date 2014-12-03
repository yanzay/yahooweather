yahooweather
============

Go implementation for getting weather from Yahoo.

Usage:
```go
package main

import (
	"fmt"
	"github.com/yanzay/yahooweather"
)

func main() {
	condition, forecasts, err := yahooweather.GetWeather(924938) // Kiev
	if err != nil {
		fmt.Printf("Error getting weather: %s", err)
		return
	}

  fmt.Printf("Current temperature: %d\n", condition.Temp)

	for _, item := range forecasts {
		fmt.Printf("%s: from %d to %d, code: %d\n", item.Date, item.Low, item.High, item.Code)
	}
}
```
