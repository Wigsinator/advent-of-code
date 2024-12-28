package main

import (
  // "fmt"
  "slices"
  "github.com/Wigsinator/advent-of-code/help"
)

type heurFunc func(directionalPos, Pos) int

type distFunc func(directionalPos, Pos) (int, bool)

type drawFunc func([]directionalPos)

func aStar(start directionalPos, goal Pos, h heurFunc, d distFunc, draw drawFunc) []directionalPos { 

  cameFrom := make(map[directionalPos]directionalPos)

  gScore := make(map[directionalPos]int)
  gScore[start] = 0

  fScore := make(map[directionalPos]int)
  fScore[start] = h(start, goal)

  openSet := help.NewHeap[directionalPos](func(a,b directionalPos) bool {return fScore[a] < fScore[b]})
  openSet.Push(start)

  for openSet.Len() > 0 {
    // fmt.Println(openSet.Visualize())
    current := openSet.Pop()
    // fmt.Println("Current is", current, "has a gScore of", gScore[current], "and an fScore of", fScore[current])
    // draw(rebuild_path(cameFrom, current, start))
    if current.Pos == goal {
      return rebuild_path(cameFrom, current, start)
    }

    for _,neighbor := range current.OldNeighbors() {
      dist, success := d(current, neighbor.Pos)
      if !success {
        continue
      }
      tentative_gScore := gScore[current] + dist
      if val, ok := gScore[neighbor]; !ok || tentative_gScore < val {
        cameFrom[neighbor] = current
        gScore[neighbor] = tentative_gScore
        fScore[neighbor] = tentative_gScore + h(neighbor, goal)
        // fmt.Println("neighbor", neighbor, "is ok. Adding with gScore of", tentative_gScore, "and an fScore of", fScore[neighbor], ". Is it in OpenSet?", openSet.Exists(neighbor))
        if !openSet.Exists(neighbor) {
          openSet.Push(neighbor)
        }
      }
    }
  }

  return nil
}

func rebuild_path(cameFrom map[directionalPos]directionalPos, current directionalPos, start directionalPos) []directionalPos {
  totalPath := make([]directionalPos,0)
  totalPath = append(totalPath, current)
  for current != start {
    current = cameFrom[current]
    totalPath = append(totalPath, current)
  }
  slices.Reverse(totalPath)
  return totalPath
}
