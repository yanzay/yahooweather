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
	forecasts, err := yahooweather.GetForecasts(924938) // Kiev
	if err != nil {
		fmt.Printf("Error getting weather: %s", err)
		return
	}
	for _, item := range forecasts {
		fmt.Printf("%s: from %d to %d, code: %d\n", item.Date, item.Low, item.High, item.Code)
	}
}
```
