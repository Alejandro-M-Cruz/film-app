package utils

import "time"

type Date struct {
    time.Time
}

func (d *Date) UnmarshalJSON(b []byte) error {
    date, err := time.Parse(time.DateOnly, string(b))
    if err != nil {
        return err
    }

    d.Time = date
    return nil
}

func (d Date) MarshalJSON() ([]byte, error) {
    return []byte(d.Time.Format(time.DateOnly)), nil
}
