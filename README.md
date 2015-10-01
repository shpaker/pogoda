[![godoc badge](http://godoc.org/github.com/shpaker/pogoda?status.png)](http://godoc.org/github.com/shpaker/pogoda)
# pogoda
Golang package for working with Yandex weather xml files.

##Installing
```bash
$ go get github.com/shpaker/pogoda
```

##Example
####Go code
```go
package main

import (
	"github.com/shpaker/pogoda"

	"flag"
	"fmt"
	"log"
)

func main() {
	var cityFlag string
	flag.StringVar(&cityFlag, "c", "Новосибирск", "Город")
	flag.Parse()
	countries, err := pogoda.NewCountries()
	if err != nil {
		log.Fatal(err)
	}
	cities, err := countries.GetCities(cityFlag)
	if err != nil {
		log.Fatal(err)
	}
	var forecast *pogoda.Fact
	pogoda.GetForecast(cities[0].Id, &forecast)
	fmt.Printf(" %s, %s, %s (%s)\n", cities[0].Country, cities[0].Part, cities[0].Name, forecast.Fact.WeatherType)
	fmt.Printf("  Температура: %g°C \n", forecast.Fact.Temperature.Value)
	fmt.Printf("  Ветер: %g м/с Направление: %s \n", forecast.Fact.WindSpeed, forecast.Fact.WindDirection)
	fmt.Printf("  Влажность: %d%% \n", forecast.Fact.Humidity)
	fmt.Printf("  Давление: %d мм рт. ст. \n", forecast.Fact.Pressure.Value)
}
```
####Running
```bash
./example -c Барнаул
```
####Result
```bash
 Россия, Алтайский край, Барнаул (ясно)
  Температура: 15°C 
  Ветер: 7 м/с Направление: sw 
  Влажность: 59% 
  Давление: 744 мм рт. ст. 

```