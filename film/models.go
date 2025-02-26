package film

import (
    "database/sql/driver"
    "errors"
    "film-app/user"
    "strconv"
    "strings"
    "time"
)

type FilmID uint

func ParseFilmID(id string) (FilmID, error) {
    parsedID, err := strconv.Atoi(id)
    if err != nil {
        return 0, errors.New("invalid film id")
    }

    return FilmID(parsedID), nil
}

type Film struct {
    ID          FilmID
    UserID      user.UserID
    User        user.User
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

func (c *Cast) Scan(value any) error {
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

func (c Cast) Value() (driver.Value, error) {
    return strings.Join(c.Members, ";"), nil
}
