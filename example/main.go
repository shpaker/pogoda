package main

import (
	"github.com/shpaker/pogoda"

	"flag"
	"fmt"
	"log"
)

var (
	cityFlag string
)

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
	pogoda.GetForecast(cities[0].Id, &forecast)

	fmt.Printf("[%d] %s, %s, %s (%s)\n", cities[0].Id, cities[0].Country, cities[0].Part, cities[0].Name, forecast.Fact.WeatherType)
	fmt.Printf("\tТемпература: %g°C \n", forecast.Fact.Temperature.Value)
	fmt.Printf("\tВетер: %g м/с \n", forecast.Fact.WindSpeed)
	fmt.Printf("\tНаправление: %s \n", forecast.Fact.WindDirection)
	fmt.Printf("\tВлажность: %d%% \n", forecast.Fact.Humidity)
	fmt.Printf("\tДавление: %d мм рт. ст. \n", forecast.Fact.Pressure.Value)
}
