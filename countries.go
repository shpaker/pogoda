package pogoda

import (
	"encoding/xml"
	"errors"
	"io/ioutil"
	"net/http"
)

type (
	Countries struct {
		Countries []*Country `xml:"country"`
	}

	Country struct {
		Cities []*City `xml:"city"`
		Name   string  `xml:"name,attr"`
	}

	City struct {
		Name    string `xml:",chardata"`
		Part    string `xml:"part,attr"`
		Country string `xml:"country,attr"`
		Id      int    `xml:"id,attr"`
		Region  int    `xml:"region,attr"`
		Head    int    `xml:"head,attr"`
		Type    int    `xml:"type,attr"`
		Resort  int    `xml:"resort,attr"`
		Climate int    `xml:"climate,attr"`
	}
)

func NewCountries() (countries *Countries, err error) {
	res, err := http.Get("https://pogoda.yandex.ru/static/cities.xml")
	if err != nil {
		return nil, err
	}
	resData, err := ioutil.ReadAll(res.Body)
	res.Body.Close()
	err = xml.Unmarshal(resData, &countries)
	if err != nil {
		return nil, err
	}
	return countries, nil
}

// Функция возвращает страну с именем "name"
func (countries *Countries) GetCountry(name string) (country *Country, err error) {
	for _, c := range countries.Countries {
		if c.Name == name {
			country = c
		}
	}
	if country == nil {
		return nil, errors.New("Country with name \"" + name + "\" not founded")
	}
	return country, nil
}

// Функция возвращает массив городов с именем "name"
func (countries *Countries) GetCities(name string) (cities []*City, err error) {
	for _, country := range countries.Countries {
		countryCities, _ := country.GetCities(name)
		cities = append(cities, countryCities...)
	}
	if len(cities) == 0 {
		err = errors.New("Cities with name \"" + name + "\" not founded")
	}
	return cities, nil
}

func (country *Country) GetCities(name string) (cities []*City, err error) {
	for _, city := range country.Cities {
		if city.Name == name {
			cities = append(cities, city)
		}
	}
	if len(cities) == 0 {
		return nil, errors.New("Cities with name \"" + name + "\" not founded")
	}
	return
}
