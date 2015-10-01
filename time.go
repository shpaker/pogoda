package pogoda

import (
	"encoding/xml"
	"regexp"
	"time"
)

// Структура данных необходимая для корректной десериализации данных о датах/времени
type Time struct {
	time.Time
}

func (t *Time) UnmarshalXML(d *xml.Decoder, start xml.StartElement) (err error) {
	var datetime string
	d.DecodeElement(&datetime, &start)
	if err = t.Unmarshal(datetime); err != nil {
		return nil
	} else {
		return err
	}
}

func (t *Time) UnmarshalText(data []byte) (err error) {
	if err = t.Unmarshal(string(data)); err != nil {
		return nil
	} else {
		return err
	}
}

func (t *Time) Unmarshal(datetime string) (err error) {
	var parse time.Time
	formats := map[string]string{
		"2006-01-02T15:04:05": "^(\\d{4}-(0|1)\\d-(0|1|2|3)\\dT(0|1|2)\\d:(0|1|2|3|4|5)\\d:(0|1|2|3|4|5)\\d)$",
		"15:04":               "^(0|1|2)\\d:(0|1|2|3|4|5)\\d)$",
		"2006-01-02":          "^(\\d{4}-(0|1)\\d-(0|1|2|3)\\d)$",
	}
	for format, reg := range formats {
		matched, _ := regexp.MatchString(reg, datetime)
		if matched == true {
			parse, err = time.Parse(format, datetime)
			if err != nil {
				return err
			}
		}
	}
	*t = Time{parse}
	return nil
}
