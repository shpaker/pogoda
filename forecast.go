// Golang пакет для работы с XML Яндекс погоды
package pogoda

import (
	"encoding/xml"
	"io/ioutil"
	"net/http"
	"strconv"
)

type (
	// Полный прогноз погоды на выбранный город.
	// Включает данные о фактическом состоянии погоды, состоянии на вчера.
	// А также данные прогноза на десять дней, включая текущий.
	Forecast struct {
		City       string           `xml:"city,attr"`      // Наименование города
		Exactname  string           `xml:"exactname,attr"` // Точное наименование
		Country    string           `xml:"country,attr"`   // Страна
		Part       string           `xml:"part,attr"`      // Регион
		Lat        string           `xml:"lat,attr"`       // Широта
		Lon        string           `xml:"lon,attr"`       // Долгота
		Id         int              `xml:"id,attr"`        // Идентификатор
		Slug       string           `xml:"slug,attr"`
		Zoom       int              `xml:"zoom,attr"`
		Source     string           `xml:"source,attr"`
		Country_id string           `xml:"country_id,attr"`
		Link       string           `xml:"link,attr"`
		Part_id    string           `xml:"part_id,attr"`
		Geoid      int              `xml:"geoid,attr"`
		Region     int              `xml:"region,attr"`
		Informer   *WeatherInformer `xml:"informer"`
		Fact       *Weather         `xml:"fact"`
		Yesterday  *Weather         `xml:"yesterday"`
		Days       []Day            `xml:"day"`
	}
	// Данные о температуре на ближайшую часть дня и следующий день.
	Informer struct {
		Informer *WeatherInformer `xml:"informer"`
	}
	// Фактические данные о погоде для выбранного города.
	Fact struct {
		Fact *Weather `xml:"fact"`
	}
	// Вчерашние данные о погоде для выбранного города.
	Yesterday struct {
		Yesterday *Weather `xml:"yesterday"`
	}
	// Прогноз на 10 дней, включая текущий.
	ByDays struct {
		Days []Day `xml:"day"`
	}
)

// "Конструктор" прогноза погоды.
// cityId - идентификатор города, s - структура для десериализации данных
func ReqForecast(cityId int, s interface{}) (err error) {
	res, err := http.Get("http://export.yandex.ru/weather-ng/forecasts/" + strconv.Itoa(cityId) + ".xml")
	if err != nil {
		return err
	}
	resData, err := ioutil.ReadAll(res.Body)
	res.Body.Close()
	err = xml.Unmarshal(resData, s)
	if err != nil {
		return err
	}
	return nil
}
