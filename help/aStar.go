package help

import (
  "slices"
)

type heuristic func(Pos) int

type distance func(Pos, Pos) (int, bool)

func AStar(start Pos, goal Pos, h heuristic, d distance) []Pos { 

  cameFrom := make(map[Pos]Pos)

  gScore := make(map[Pos]int)
  gScore[start] = 0

  fScore := make(map[Pos]int)
  fScore[start] = h(start)

  openSet := NewHeap[Pos](func(a,b Pos) bool {return fScore[a] < fScore[b]})
  openSet.Push(start)

  for openSet.Len() > 0 {
    current := openSet.Pop()
    if current == goal {
      return rebuild_path(cameFrom, current, start)
    }

    for _,neighbor := range current.Neighbors() {
      dist, success := d(current, neighbor)
      if !success {
        continue
      }
      tentative_gScore := gScore[current] + dist
      if val, ok := gScore[neighbor]; !ok || tentative_gScore < val {
        cameFrom[neighbor] = current
        gScore[neighbor] = tentative_gScore
        fScore[neighbor] = tentative_gScore + h(neighbor)
        if !openSet.Exists(neighbor) {
          openSet.Push(neighbor)
        }
      }
    }
  }

  return nil
}

func rebuild_path(cameFrom map[Pos]Pos, current Pos, start Pos) []Pos {
  totalPath := make([]Pos,0)
  totalPath = append(totalPath, current)
  for current != start {
    current = cameFrom[current]
    totalPath = append(totalPath, current)
  }
  slices.Reverse(totalPath)
  return totalPath
}
