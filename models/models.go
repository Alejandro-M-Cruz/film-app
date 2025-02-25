package models

import (
    "database/sql/driver"
    "errors"
    "strings"
    "time"
)

type UserID = uint

type User struct {
    ID        UserID
    Username  string
    Password  string
    CreatedAt time.Time
}

type FilmID = uint

type Film struct {
    ID          FilmID
    UserID      UserID
    User        User
    Title       string
    ReleaseDate time.Time
    Genre       string
    Director    string
    Cast
    Synopsis  string
    CreatedAt time.Time
    UpdatedAt time.Time
}

type Cast struct {
    Members []string
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
