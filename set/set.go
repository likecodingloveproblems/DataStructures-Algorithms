package set

type ISet[T comparable] interface {
	Add(T)
	Remove(T) error
	Contains(T) bool
	Clear()
}

type Set[T comparable] struct {
	items map[T]struct{}
}

func NewSet[T comparable]() *Set[T] {
	return &Set[T]{items: map[T]struct{}{}}
}

func (s *Set[T]) Add(item T) {
	s.items[item] = struct{}{}
}

func (s *Set[T]) Remove(item T) error {
	if s.IsEmpty() {
		return ErrIsEmpty
	}
	delete(s.items, item)
	return nil
}

func (s *Set[T]) Contains(item T) bool {
	_, e := s.items[item]
	return e
}

func (s *Set[T]) Clear() {
	s.items = map[T]struct{}{}
}

func (s *Set[T]) IsEmpty() bool {
	return len(s.items) == 0
}
