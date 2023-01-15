package stack

type Stack[T any] struct {
	count uint
	items []T
}

func (s *Stack[T]) GetCount() uint {
	return s.count
}

func (s *Stack[T]) ToArray() []T {
	if s.IsEmpty() {
		return []T{}
	}
	return s.items[:s.count]
}

func (s *Stack[T]) Push(item T) {
	if s.count < uint(len(s.items)) {
		s.items[s.count] = item
	} else {
		s.items = append(s.items, item)
	}
	s.count++
}

func (s *Stack[T]) Pop() (T, error) {
	if s.IsEmpty() {
		var noop T
		return noop, StackIsEmpty
	}
	s.count--
	return s.items[s.count], nil
}

func (s *Stack[T]) Peek() (T, error) {
	if s.IsEmpty() {
		var noop T
		return noop, StackIsEmpty
	}
	return s.items[s.count-1], nil
}

func (s *Stack[T]) IsEmpty() bool {
	return s.count == 0
}
