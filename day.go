package pogoda

// Структура Day описывает данные для десериализации прогноза на день
// Яндекс предоставляет данные на 10 дней включая текущий
// В состав прогноза входят данные о погоде в разные части дня.
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
		// Прогноз на части дня (1 утро, 2 день, 3 вечер, 4 ночь)
		DayParts []struct {
			TypeId          int     `xml:"typeid,attr"`
			Type            string  `xml:"type,attr"`
			TemperatureFrom float32 `xml:"temperature_from"`
			TemperatureTo   float32 `xml:"temperature_to"`
			TemperatureData struct {
				Avg struct {
					Value   int    `xml:",chardata"`
					Bgcolor string `xml:"bgcolor,attr"`
				} `xml:"avg"`
				From int `xml:"from"`
				To   int `xml:"to"`
			} `xml:"temperature-data"`
			WeatherCondition *Condition `xml:"weather_condition"`
			Image            *Image     `xml:"image"`
			ImageV2          *ImageV2   `xml:"image-v2"`
			ImageV3          *ImageV3   `xml:"image-v3"`
			WeatherType      string     `xml:"weather_type"`    // Тип погоды на русском
			WeatherTypeTt    string     `xml:"weather_type_tt"` // Тип погоды на (?)
			WeatherTypeTr    string     `xml:"weather_type_tr"` // Тип погоды на турецком(?)
			WeatherTypeKz    string     `xml:"weather_type_kz"` // Тип погоды на казахском
			WeatherTypeUa    string     `xml:"weather_type_ua"` // Тип погоды на украинском
			WeatherTypeBy    string     `xml:"weather_type_by"` // Тип погоды на белорусском
			WindDirection    string     `xml:"wind_direction"`  // Направление ветра
			WindSpeed        float32    `xml:"wind_speed"`      // Скорость ветра
			Humidity         int        `xml:"humidity"`        // Влажность воздуха
			Pressure         *Pressure  `xml:"pressure"`        // Атмосферное давление
			MslpPressure     *Pressure  `xml:"mslp_pressure"`   // Давление на уровне моря
		} `xml:"day_part"`
		// Почасовой прогноз
		Hours []struct {
			At               int        `xml:"at,attr"`           // Час
			Temperature      float32    `xml:"temperature"`       // Прогнозируемая температура
			WeatherCondition *Condition `xml:"weather_condition"` // Состояние погоды
			Image            *Image     `xml:"image"`
			ImageV2          *ImageV2   `xml:"image-v2"`
			ImageV3          *ImageV3   `xml:"image-v3"`
		} `xml:"hour"`
	}
)
