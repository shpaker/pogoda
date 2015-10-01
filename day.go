package pogoda

// Структура Day описывает данные для десериализации прогноза на день
// Яндекс предоставляет данные на 10 дней включая текущий
type (
	Day struct {
		Date      *Time `xml:"date,attr"` // Дата прогноза
		Sunrise   *Time `xml:"sunrise"`   // Время восхода солнца
		Sunset    *Time `xml:"sunset"`    // Время заката солнца
		Moonrise  *Time `xml:"moonrise"`  // Время появления луны
		Moonset   *Time `xml:"moonset"`   // Время исчезновения луны
		MoonPhase *struct {
			Value int    `xml:",chardata"`
			Code  string `xml:"code,attr"`
		} `xml:"moon_phase"`
		Biomet *struct {
			Index    int `xml:"index,attr"`
			Geomag   int `xml:"geomag,attr"`
			Uv       int `xml:"uv,attr"`
			Messages []struct {
				Code string `xml:"code,attr"`
			} `xml:"message"`
		} `xml:"biomet"`
		DayParts []*DayPart `xml:"day_part"` // Прогноз на части дня (утро, деньб вечер, ночь)
		Hours    []*Hour    `xml:"hour"`     // Почасовой прогноз
	}
)
