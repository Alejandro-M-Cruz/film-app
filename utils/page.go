package utils

type Page[T any] struct {
    Data       []T
    Page       int
    PageSize   int
    TotalPages int
}
