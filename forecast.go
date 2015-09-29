package pogoda

import (
	"encoding/xml"
	"io/ioutil"
	"net/http"
	"strconv"
)

type (
	Forecast struct {
		City       string `xml:"city,attr"`
		Exactname  string `xml:"exactname,attr"`
		Slug       string `xml:"slug,attr"`
		Country    string `xml:"country,attr"`
		Part       string `xml:"part,attr"`
		Lat        string `xml:"lat,attr"`
		Lon        string `xml:"lon,attr"`
		Zoom       int    `xml:"zoom,attr"`
		Id         int    `xml:"id,attr"`
		Source     string `xml:"source,attr"`
		Country_id string `xml:"country_id,attr"`
		Link       string `xml:"link,attr"`
		Part_id    string `xml:"part_id,attr"`
		Geoid      int    `xml:"geoid,attr"`
		Region     int    `xml:"region,attr"`
		Informer
		Fact
		Yesterday
		ByDays
	}
	Informer struct {
		Informer *WeatherInformer `xml:"informer"`
	}
	Fact struct {
		Fact *Weather `xml:"fact"`
	}
	Yesterday struct {
		Yesterday *Weather `xml:"yesterday"`
	}
	ByDays struct {
		Days []Day `xml:"day"`
	}
)

func GetForecast(cityId int, s interface{}) (err error) {
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
