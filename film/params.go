package film

import (
    "strconv"
    "time"
)

type IndexParams struct {
    Page     int
    PageSize int
    Filters  []Filter
}

var PageSizeOptions = map[int]bool{
    10: true,
    20: true,
    50: true,
}

const DefaultPageSize = 10

func NewIndexParams(params map[string][]string) IndexParams {
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

    return IndexParams{
        Page:     page,
        PageSize: pageSize,
        Filters:  ParseFilters(params),
    }
}

type Filter interface {
    Name() string
}

func ParseFilters(params map[string][]string) []Filter {
    filters := make([]Filter, 0)

    for key, values := range params {
        switch key {
        case "title":
            filters = append(filters, FilterByTitle{PartialTitle: values[0]})
        case "genres":
            filters = append(filters, FilterByGenres{Genres: values})
        case "released_after":
            date, err := time.Parse(time.RFC3339, values[0])
            if err != nil {
                continue
            }
            filters = append(filters, FilterByReleaseDateAfter{Date: date})
        case "released_before":
            date, err := time.Parse(time.RFC3339, values[0])
            if err != nil {
                continue
            }
            filters = append(filters, FilterByReleaseDateBefore{Date: date})
        case "released_between":
            start, err := time.Parse(time.RFC3339, values[0])
            if err != nil {
                continue
            }
            end, err := time.Parse(time.RFC3339, values[1])
            if err != nil {
                continue
            }
            filters = append(filters, FilterByReleaseDateBetween{Start: start, End: end})
        }
    }

    return filters
}

type FilterByTitle struct {
    PartialTitle string
}

func (f FilterByTitle) Name() string {
    return "title"
}

type FilterByGenres struct {
    Genres []string
}

func (f FilterByGenres) Name() string {
    return "genres"
}

type FilterByReleaseDateAfter struct {
    Date time.Time
}

func (f FilterByReleaseDateAfter) Name() string {
    return "released_after"
}

type FilterByReleaseDateBefore struct {
    Date time.Time
}

func (f FilterByReleaseDateBefore) Name() string {
    return "released_before"
}

type FilterByReleaseDateBetween struct {
    Start time.Time
    End   time.Time
}

func (f FilterByReleaseDateBetween) Name() string {
    return "released_between"
}
