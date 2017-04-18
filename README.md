[![godoc badge](http://godoc.org/github.com/shpaker/pogoda?status.png)](http://godoc.org/github.com/shpaker/pogoda)
# pogoda
Golang package for working with Yandex weather xml files.

## Installing
```bash
$ go get github.com/shpaker/pogoda
```

## Example
#### Go code
```go
package main

import (
	"github.com/shpaker/pogoda"

	"flag"
	"fmt"
	"log"
)

var cityFlag string

func init() {
	flag.StringVar(&cityFlag, "c", "Новосибирск", "Город")
	flag.Parse()
}

func main() {
	countries, err := pogoda.NewCountries()
	if err != nil {
		log.Fatal(err)
	}
	cities, err := countries.GetCities(cityFlag)
	if err != nil {
		log.Fatal(err)
	}
	var forecast *pogoda.Forecast
	pogoda.ReqForecast(cities[0].Id, &forecast)
	fmt.Printf("[%d] %s, %s, %s (%s)\n", cities[0].Id, cities[0].Country, cities[0].Part, cities[0].Name, forecast.Fact.WeatherType)
	fmt.Printf("\tТемпература: %g°C \n", forecast.Fact.Temperature.Value)
	fmt.Printf("\tВетер: %g м/с \n", forecast.Fact.WindSpeed)
	fmt.Printf("\tНаправление: %s \n", forecast.Fact.WindDirection)
	fmt.Printf("\tВлажность: %d%% \n", forecast.Fact.Humidity)
	fmt.Printf("\tДавление: %d мм рт. ст. \n", forecast.Fact.Pressure.Value)
}

```
#### Running
```bash
./example -c Барнаул
```
#### Result
```bash
 Россия, Алтайский край, Барнаул (ясно)
  Температура: 15°C 
  Ветер: 7 м/с Направление: sw 
  Влажность: 59% 
  Давление: 744 мм рт. ст. 

```
