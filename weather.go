package pogoda

type (
	// Данные о погоде для выбранного города.
	// Предоставляеются фактические данные <fact> и данные на вчерашний день <yesterday>
	Weather struct {
		Station []struct {
			Value    string `xml:",chardata"`
			Lang     string `xml:"lang,attr"`
			Distance int    `xml:"distance,attr"`
		} `xml:"station"`
		ObservationTime *Time     `xml:"observation_time"` // Время наблюдения данных
		Uptime          *Time     `xml:"uptime"`
		Temperature     *struct { // Данные температуры
			Value float32 `xml:",chardata"` // Значение температуры
			Color string  `xml:"color,attr"`
			Plate string  `xml:"plate,attr"`
		} `xml:"temperature"`
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
		Daytime          string     `xml:"daytime"`         // Часть дня (m|d|e|n)
		Season           *Season    `xml:"season"`          // Время года (предоставляется на англ.)
		IpadImage        string     `xml:"ipad_image"`
	}

	// Состояние погоды (Например: малооблочно, дождливо и пр.)
	Condition struct {
		Code string `xml:"code,attr"`
	}

	// Атмосферное давление
	Pressure struct {
		Value int    `xml:",chardata"`
		Units string `xml:"units,attr"`
	}
	// Время года (предоставляется на англ.)
	Season struct {
		Value string `xml:",chardata"`
		Type  string `xml:"type"`
	}

	Image struct {
		Value string `xml:",chardata"`
		Type  int    `xml:"type,attr"`
	}

	ImageV2 struct {
		Value string `xml:",chardata"`
		Type  int    `xml:"type,attr"`
		Color string `xml:"color,attr"`
	}

	ImageV3 struct {
		Value string `xml:",chardata"`
		Type  int    `xml:"type,attr"`
	}
)
