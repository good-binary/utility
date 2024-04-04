package utils

type Slicer[T any] struct {
	slice []T
}

func NewSlicer[T any](slice []T) *Slicer[T] {
	return &Slicer[T]{slice}
}

func (s *Slicer[T]) Len() int {
	return len(s.slice)
}

func (s *Slicer[T]) Get(i int) T {
	return s.slice[i]
}

func (s *Slicer[T]) Set(i int, value T) {
	s.slice[i] = value
}

func (s *Slicer[T]) Append(value T) {
	s.slice = append(s.slice, value)
}

func (s *Slicer[T]) Prepend(value T) {
	s.slice = append([]T{value}, s.slice...)
}

func (s *Slicer[T]) Remove(i int) {
	s.slice = append(s.slice[:i], s.slice[i+1:]...)
}

func (s *Slicer[T]) Clear() {
	s.slice = []T{}
}
