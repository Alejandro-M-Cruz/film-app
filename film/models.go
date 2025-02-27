package film

import (
    "database/sql/driver"
    "errors"
    "film-app/user"
    "film-app/utils"
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
    ReleaseDate utils.Date
    Genre       Genre
    Director    string
    Cast
    Synopsis  string
    CreatedAt time.Time
    UpdatedAt time.Time
}

type Genre string

func (g *Genre) Scan(value any) error {
    genre, ok := value.(string)
    if !ok {
        return errors.New("genre must be a string")
    }

    *g = Genre(genre)
    return nil
}

func (g Genre) Value() (driver.Value, error) {
    return string(g), nil
}

const (
    Action    Genre = "action"
    Adventure Genre = "adventure"
    Comedy    Genre = "comedy"
    Drama     Genre = "drama"
    Fantasy   Genre = "fantasy"
    Horror    Genre = "horror"
    Mystery   Genre = "mystery"
    Romance   Genre = "romance"
    Thriller  Genre = "thriller"
    Western   Genre = "western"
)

var Genres = map[string]Genre{
    "action":    Action,
    "adventure": Adventure,
    "comedy":    Comedy,
    "drama":     Drama,
    "fantasy":   Fantasy,
    "horror":    Horror,
    "mystery":   Mystery,
    "romance":   Romance,
    "thriller":  Thriller,
    "western":   Western,
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
