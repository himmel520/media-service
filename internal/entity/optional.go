package entity

type Optional[T any] struct {
	Value T
	Set   bool
}
