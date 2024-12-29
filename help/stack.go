package help

type Stack[T any] struct {
  data []T
}

func NewStack[T any]() (*Stack[T]) {
  return &Stack[T]{nil}
}

func (s *Stack[T]) Push(v T) {
  s.data = append(s.data, v)
}

func (s *Stack[T]) Pop() T {
  v := s.data[len(s.data)-1]
  s.data = s.data[:len(s.data)-1]
  return v
}

func (s *Stack[T]) IsEmpty() bool {
  return len(s.data) == 0
}
