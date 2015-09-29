package pogoda

type (
	DayPart struct {
		TypeId           int              `xml:"typeid,attr"`
		Type             string           `xml:"type,attr"`
		TemperatureFrom  int              `xml:"temperature_from"`
		TemperatureTo    int              `xml:"temperature_to"`
		TemperatureData  *TemperatureData `xml:"temperature-data"`
		WeatherCondition *Condition       `xml:"weather_condition"`
		Image            *Image           `xml:"image"`
		ImageV2          *ImageV2         `xml:"image-v2"`
		ImageV3          *ImageV3         `xml:"image-v3"`
		WeatherType      string           `xml:"weather_type"`
		WeatherTypeTt    string           `xml:"weather_type_tt"`
		WeatherTypeTr    string           `xml:"weather_type_tr"`
		WeatherTypeKz    string           `xml:"weather_type_kz"`
		WeatherTypeUa    string           `xml:"weather_type_ua"`
		WeatherTypeBy    string           `xml:"weather_type_by"`
		WinDirection     string           `xml:"wind_direction"`
		WindSpeed        float32          `xml:"wind_speed"`
		Humadity         int              `xml:"humadity"`
		Pressure         *Pressure        `xml:"pressure"`
		MslpPressure     *Pressure        `xml:"mslp_pressure"`
	}

	TemperatureData struct {
		Avg  *Avg `xml:"avg"`
		From int  `xml:"from"`
		To   int  `xml:"to"`
	}

	Avg struct {
		Value   int    `xml:",chardata"`
		Bgcolor string `xml:"bgcolor,attr"`
	}
)
