package help

type Queue[T any] struct {
  data []T
}

func NewQueue[T any]() (*Queue[T]) {
  return &Queue[T]{data:make([]T,0)}
}

func (q *Queue[T]) Push(v T) {
  q.data = append(q.data,v)
}

func (q *Queue[T]) Pop() T {
  v := q.data[0]
  q.data = q.data[1:]
  return v
}

func (q *Queue[T]) IsEmpty() bool {
  return len(q.data) == 0
}
