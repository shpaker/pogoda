package pogoda

type Hour struct {
	At               int        `xml:"at,attr"`
	Temperature      int        `xml:"temperature"`
	WeatherCondition *Condition `xml:"weather_condition"`
	Image            *Image     `xml:"image"`
	ImageV2          *ImageV2   `xml:"image-v2"`
	ImageV3          *ImageV3   `xml:"image-v3"`
}
