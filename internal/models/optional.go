package models

type Optional[T any] struct {
	Value T
	Set   bool
}
