package utils

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"time"
)

type Date struct {
	time.Time
}

func NewDate(date time.Time) Date {
	return Date{date}
}

func ParseDate(value string) (Date, error) {
	t, err := time.Parse(time.DateOnly, value)
	return NewDate(t), err
}

func (d *Date) Scan(value any) error {
	date, ok := value.(time.Time)
	if !ok {
		return errors.New("date must be a time.Time")
	}

	d.Time = date
	return nil
}

func (d Date) Value() (driver.Value, error) {
	return d.Time, nil
}

func (d *Date) UnmarshalJSON(b []byte) error {
	var dateStr string
	if err := json.Unmarshal(b, &dateStr); err != nil {
		return err
	}

	date, err := ParseDate(dateStr)
	if err != nil {
		return err
	}

	*d = date
	return err
}

func (d Date) MarshalJSON() ([]byte, error) {
	return json.Marshal(d.Time.Format(time.DateOnly))
}
