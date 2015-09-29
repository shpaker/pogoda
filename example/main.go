package main

import (
	"fmt"
	"log"
	"pogoda"
)

func main() {
	countries, err := pogoda.NewCountries()
	if err != nil {
		log.Fatal(err)
	}
	cities, err := countries.GetCities("Барнаул")
	if err != nil {
		log.Fatal(err)
	}
	var forecast *pogoda.Fact
	pogoda.GetForecast(cities[0].Id, &forecast)
	fmt.Printf("%s, %s, %s (%s)\n", cities[0].Country, cities[0].Part, cities[0].Name, forecast.Fact.WeatherType)
	fmt.Printf("Температура: %g°C \n", forecast.Fact.Temperature.Value)
	fmt.Printf("Ветер: %g м/с Направление: %s \n", forecast.Fact.WindSpeed, forecast.Fact.WindDirection)
	fmt.Printf("Влажность: %d%% \n", forecast.Fact.Humidity)
	fmt.Printf("Давление: %d мм рт. ст.", forecast.Fact.Pressure.Value)
}
