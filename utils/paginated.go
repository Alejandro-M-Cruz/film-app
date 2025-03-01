package utils

type Paginated[T any] struct {
	Items    []T
	Total    int64
	Page     int
	PageSize int
}
