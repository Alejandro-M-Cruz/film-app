package utils

import (
    "encoding/json"
    "time"
)

type Date struct {
    time.Time
}

func NewDate(date time.Time) Date {
    return Date{date}
}

func (d *Date) UnmarshalJSON(b []byte) error {
    var dateStr string
    if err := json.Unmarshal(b, &dateStr); err != nil {
        return err
    }

    date, err := time.Parse(time.DateOnly, dateStr)
    if err != nil {
        return err
    }

    d.Time = date
    return nil
}

func (d Date) MarshalJSON() ([]byte, error) {
    return json.Marshal(d.Time.Format(time.DateOnly))
}
