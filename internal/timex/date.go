package timex

import (
	"encoding/json"
	"fmt"
	"time"
)

type Date struct {
	t time.Time
}

func NewDateFromTime(t time.Time) Date {
	return Date{t: t}
}

func NewDateFromString(s string) (Date, error) {
	t, err := time.Parse(time.DateOnly, s)
	if err != nil {
		return Date{}, fmt.Errorf("failed to parse date: %w", err)
	}
	return Date{t: t}, nil
}

func (d *Date) UnmarshalJSON(b []byte) error {
	var str string
	if err := json.Unmarshal(b, &str); err != nil {
		return fmt.Errorf("(*Date) cannot unmarshal: failed to unmarshal string: %w", err)
	}

	date, err := NewDateFromString(str)
	if err != nil {
		return fmt.Errorf("(*Date) cannot unmarshal: failed to unmarshal string: %w", err)
	}

	*d = date
	return nil
}

func (d *Date) MarshalJSON() ([]byte, error) {
	return json.Marshal(d.String())
}

func (d Date) String() string {
	return d.t.Format(time.DateOnly)
}

func (d Date) Time() time.Time {
	return d.t
}
