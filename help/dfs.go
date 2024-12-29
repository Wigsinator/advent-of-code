package help

import "fmt"

type adjacent[T comparable] func(T) []T 
type validate[T comparable] func(T) bool

func Dfs[T comparable](vertex T, adjacentFunc adjacent[T], validateFunc validate[T]) []T {
  DEBUG := true
  if DEBUG {fmt.Println("Initiating DFS")}
  s := NewStack[T]()
  discovered := make(map[T]bool)
  s.Push(vertex)
  successes := make([]T, 0)

  for !s.IsEmpty() {
    if DEBUG {fmt.Println("Current Stack State:", s)}
    v := s.Pop()
    if !discovered[v] {
      if validateFunc(v) { successes = append(successes, v) }
      discovered[v] = true
      for _,w := range adjacentFunc(v){
        s.Push(w)
      }
    }
  }
  return successes
}
