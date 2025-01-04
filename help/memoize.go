package help

type Memoized[T comparable, V any] struct {
  f func(T) (V)
  cache map[T]V
}

func Memoize[T comparable, V any](f func(T) V) *Memoized[T,V] {
  return &Memoized[T,V]{f: f, cache: make(map[T]V)}
}

func (m *Memoized[T, V]) Call(x T) V {
  if val, ok :=  m.cache[x]; ok {
    return val
  }
  result := m.f(x)
  m.cache[x] = result
  return result
}
