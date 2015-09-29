package pogoda

type (
	WeatherInformer struct {
		Temperature []InformerTemperature `xml:"temperature"`
	}

	InformerTemperature struct {
		Value float32 `xml:",chardata"`
		Color string  `xml:"color,attr"`
		Type  string  `xml:"type,attr"`
	}
)
