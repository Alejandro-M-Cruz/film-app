package models

import (
    "database/sql/driver"
    "errors"
    "strings"
    "time"
)

type UserID = uint

type User struct {
    ID        UserID    `json:"id"`
    Username  string    `json:"username"`
    Password  string    `json:"-"`
    CreatedAt time.Time `json:"created_at"`
}

type FilmID = uint

type Film struct {
    ID          FilmID    `json:"id"`
    UserID      UserID    `json:"user_id"`
    User        User      `json:"user"`
    Title       string    `json:"title"`
    ReleaseDate time.Time `json:"release_date"`
    Genre       string    `json:"genre"`
    Director    string    `json:"director"`
    Cast
    Synopsis  string    `json:"synopsis"`
    CreatedAt time.Time `json:"created_at"`
    UpdatedAt time.Time `json:"updated_at"`
}

type Cast struct {
    Members []string `json:"cast"`
}

func (c *Cast) Scan(value interface{}) error {
    membersStr, ok := value.(string)
    if !ok {
        return errors.New("cast must be a string")
    }

    members := strings.Split(membersStr, ";")

    for _, member := range members {
        c.Members = append(c.Members, member)
    }

    return nil
}

func (c *Cast) Value() (driver.Value, error) {
    return strings.Join(c.Members, ";"), nil
}
