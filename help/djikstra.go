package help

// Dist func should return -1 on fail
type dDistance[T comparable] func(T, T) int

type neighbors[T comparable] func(T) []T

func djikstra[T comparable](start T, goal T, distFunc dDistance[T], neighborFunc neighbors[T], goalSeeking bool) (map[T]int, map[T]T) {
  distMap := make(map[T]int)
  prevMap := make(map[T]T)

  openSet := NewHeap[T](func(a,b T) bool {return distMap[a]< distMap[b]})
  distMap[start] = 0
  openSet.Push(start)

  for openSet.Len() > 0 {
    current := openSet.Pop()
    if goalSeeking && current == goal {
      break
    }

    for _, neighbor := range neighborFunc(current) {
      dist := distFunc(current, neighbor)
      if dist < 0 { continue }
      tempScore := distMap[current] + dist
      if val, ok := distMap[neighbor]; !ok || tempScore < val {
        prevMap[neighbor] = current
        distMap[neighbor] = tempScore
        if !openSet.Exists(neighbor) {
          openSet.Push(neighbor)
        }
      }
    }
  }

  return distMap, prevMap
}

func Djikstra[T comparable](start T, goal T, distFunc dDistance[T], neighborFunc neighbors[T]) (map[T]int, map[T]T) {
  return djikstra[T](start, goal, distFunc, neighborFunc, true)
}

func DjikstraFull[T comparable](start T, distFunc dDistance[T], neighborFunc neighbors[T]) (map[T]int, map[T]T) {
  return djikstra[T](start, start, distFunc, neighborFunc, false)
}

func DjikstraPos(start Pos, goal Pos, distFunc dDistance[Pos]) (map[Pos]int, map[Pos]Pos) {
  return djikstra[Pos](start, goal, distFunc, func(p Pos) []Pos {return p.Neighbors()}, true)
}

func DjikstraPosFull(start Pos, goal Pos, distFunc dDistance[Pos]) (map[Pos]int, map[Pos]Pos) {
  return djikstra[Pos](start, start, distFunc, func(p Pos) []Pos {return p.Neighbors()}, false)
}
