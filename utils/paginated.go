package utils

type Paginated[T any] struct {
    Data       []T
    Page       int
    PageSize   int
    TotalPages int
}
