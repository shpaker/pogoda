package pogoda

type (
	Weather struct {
		Station          []Station    `xml:"station"`
		ObservationTime  *Time        `xml:"observation_time"`
		Uptime           *Time        `xml:"uptime"`
		Temperature      *Temperature `xml:"temperature"`
		WeatherCondition *Condition   `xml:"weather_condition"`
		Image            *Image       `xml:"image"`
		ImageV2          *ImageV2     `xml:"image-v2"`
		ImageV3          *ImageV3     `xml:"image-v3"`
		WeatherType      string       `xml:"weather_type"`
		WeatherTypeTt    string       `xml:"weather_type_tt"`
		WeatherTypeTr    string       `xml:"weather_type_tr"`
		WeatherTypeKz    string       `xml:"weather_type_kz"`
		WeatherTypeUa    string       `xml:"weather_type_ua"`
		WeatherTypeBy    string       `xml:"weather_type_by"`
		WindDirection    string       `xml:"wind_direction"`
		WindSpeed        float32      `xml:"wind_speed"`
		Humidity         int          `xml:"humidity"`
		Pressure         *Pressure    `xml:"pressure"`
		MslpPressure     *Pressure    `xml:"mslp_pressure"`
		Daytime          string       `xml:"daytime"`
		Season           *Season      `xml:"season"`
		IpadImage        string       `xml:"ipad_image"`
	}

	Station struct {
		Value    string `xml:",chardata"`
		Lang     string `xml:"lang,attr"`
		Distance int    `xml:"distance,attr"`
	}

	Temperature struct {
		Value float32 `xml:",chardata"`
		Color string  `xml:"color,attr"`
		Plate string  `xml:"plate,attr"`
	}

	Condition struct {
		Code string `xml:"code,attr"`
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

	Pressure struct {
		Value int    `xml:",chardata"`
		Units string `xml:"units,attr"`
	}

	Season struct {
		Value string `xml:",chardata"`
		Type  string `xml:"type"`
	}
)
