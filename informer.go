package pogoda

// Данные о температуре на ближайшую часть дня и следующий день
type WeatherInformer struct {
	Temperature []struct {
		Value float32 `xml:",chardata"`
		Color string  `xml:"color,attr"`
		Type  string  `xml:"type,attr"`
	} `xml:"temperature"`
}
