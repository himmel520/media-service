package paging

import (
	"math"

	"golang.org/x/exp/constraints"
)

// CalculatePages вычисляет количество страниц для пагинации
func CalculatePages[T1, T2 constraints.Float | constraints.Integer](total T1, perPage T2) uint64 {
	if perPage <= 0 {
		return 0
	}

	return uint64(math.Ceil(float64(total) / float64(perPage)))
}
