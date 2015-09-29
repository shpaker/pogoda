package pogoda

type (
	Day struct {
		Date      *Time      `xml:"date,attr"`
		Sunrise   *Time      `xml:"sunrise"`
		Sunset    *Time      `xml:"sunset"`
		MoonPhase *MoonPhase `xml:"moon_phase"`
		Moonrise  *Time      `xml:"moonrise"`
		Moonset   *Time      `xml:"moonset"`
		Biomet    *Biomet    `xml:"biomet"`
		DayParts  []*DayPart `xml:"day_part"`
		Hours     []*Hour    `xml:"hour"`
	}

	MoonPhase struct {
		Value int    `xml:",chardata"`
		Code  string `xml:"code,attr"`
	}

	Biomet struct {
		Index    int             `xml:"index,attr"`
		Geomag   int             `xml:"geomag,attr"`
		Uv       int             `xml:"uv,attr"`
		Messages []BiometMessage `xml:"message"`
	}

	BiometMessage struct {
		Code string `xml:"code,attr"`
	}
)
