package film

import (
    "strconv"
    "strings"
    "time"
)

type Params struct {
    Page     int     `json:"page"`
    PageSize int     `json:"page_size"`
    Filters  Filters `json:"filters"`
}

var PageSizeOptions = map[int]bool{
    10: true,
    20: true,
    50: true,
}

const DefaultPageSize = 10

func NewParams(params map[string][]string) Params {
    pageStr, ok := params["page"]
    page := 1
    var err error

    if ok {
        page, err = strconv.Atoi(pageStr[0])
        if err != nil {
            page = 1
        }
    }

    pageSizeStr, ok := params["page_size"]
    pageSize := DefaultPageSize

    if ok {
        pageSize, err = strconv.Atoi(pageSizeStr[0])
        if err != nil || !PageSizeOptions[pageSize] {
            pageSize = DefaultPageSize
        }
    }

    return Params{
        Page:     page,
        PageSize: pageSize,
        Filters:  ParseFilters(params),
    }
}

type Filters struct {
    Title          string     `json:"title"`
    Genres         []Genre    `json:"genres"`
    ReleasedAfter  *time.Time `json:"released_after"`
    ReleasedBefore *time.Time `json:"released_before"`
}

func ParseFilters(params map[string][]string) Filters {
    return Filters{
        Title:          parseTitleFilter(params),
        Genres:         parseGenresFilter(params),
        ReleasedAfter:  parseReleaseDateAfterFilter(params),
        ReleasedBefore: parseReleaseDateBeforeFilter(params),
    }
}

func parseTitleFilter(params map[string][]string) string {
    return extractParam(params, "title")
}

func parseGenresFilter(params map[string][]string) []Genre {
    genreStrings := strings.Split(extractParam(params, "genres"), ",")
    genres := make([]Genre, 0)

    for _, genre := range genreStrings {
        if _, ok := Genres[genre]; !ok {
            continue
        }
        genres = append(genres, Genre(genre))
    }

    return genres
}

func parseReleaseDateAfterFilter(params map[string][]string) *time.Time {
    dateStr := extractParam(params, "released_after")

    if dateStr == "" {
        return nil
    }

    date, err := time.Parse(time.DateOnly, dateStr)
    if err != nil {
        return nil
    }

    return &date
}

func parseReleaseDateBeforeFilter(params map[string][]string) *time.Time {
    dateStr := extractParam(params, "released_before")

    if dateStr == "" {
        return nil
    }

    date, err := time.Parse(time.DateOnly, dateStr)
    if err != nil {
        return nil
    }

    return &date
}

func extractParam(params map[string][]string, key string) string {
    values, ok := params[key]
    if !ok || len(values) == 0 {
        return ""
    }

    return values[0]
}
