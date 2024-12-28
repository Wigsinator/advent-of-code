package help

import "fmt"

// Create Heaps such that cmp(parent, child) must be true
type Heap[T comparable] struct {
  data []T
  comp func(a,b T) bool
}

func NewHeap[T comparable](comp func(a,b T) bool) *Heap[T] {
  return &Heap[T]{comp: comp}
}

func (h *Heap[T]) Len() int { return len(h.data) }

// Place an element at the end of the heap, then send it up until it is in a valid placement.
func (h *Heap[T]) Push(v T) {
  h.data = append(h.data, v)
  h.up(h.Len()-1)
}

// Swap the root element to the end, 
func (h *Heap[T]) Pop() T {
  n := h.Len() - 1
  if n > 0 {
    h.swap(0,n)
    h.down()
  }
  v := h.data[n]
  h.data = h.data[:n]
  return v
}

func (h *Heap[T]) Exists(v T) bool {
  for _, el := range h.data {
    if el == v { return true }
  }
  return false
}

func (h *Heap[T]) Visualize() (output string) {
  for i := range h.Len() {
    output += fmt.Sprint(h.data[i],", ")
  }
  return output
}

func (h *Heap[T]) swap(i, j int) {
  h.data[i], h.data[j] = h.data[j], h.data[i]
}

// Send element jj up the heap until cmp(child, parent) is false.
func (h *Heap[T]) up(jj int) {
  for {
    i := parent(jj)
    if i == jj || !h.comp(h.data[jj], h.data[i]) {
      break
    }
    h.swap(i, jj)
    jj = i
  }
}

// Send the root element down the heap until cmp(child, parent) is false. This ignores the last element in the heap, because down is only called when Popping
func (h *Heap[T]) down() {
  n := h.Len() - 1
  i := 0
  for {
    j1 := left(i)
    if j1 >= n || j1 < 0 {
      break
    }
    j := j1
    j2 := right(i)
    if j2 < n && h.comp(h.data[j2], h.data[j1]) {
      j = j2
    }
    if !h.comp(h.data[j], h.data[i]) {
      break
    }
    h.swap(i, j)
    i = j
  }
}

func parent(i int) int { return (i - 1) / 2}
func left(i int) int { return (i * 2) + 1 }
func right(i int) int { return left(i) + 1 }
