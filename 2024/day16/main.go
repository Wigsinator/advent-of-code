package main

import (
  _ "embed"
  "flag"
  "fmt"
  "strings"

  "github.com/Wigsinator/advent-of-code/util"
  "github.com/Wigsinator/advent-of-code/help"
)

//go:embed input.txt
var input string

func init() {
  // do this in init (not main) so test file has same input
  input = strings.TrimRight(input, "\n")
  if len(input) == 0 {
    panic("empty input.txt file")
  }
}

func main() {
  var part int
  flag.IntVar(&part, "part", 1, "part 1 or 2")
  flag.Parse()
  fmt.Println("Running part", part)

  if part == 1 {
    ans := part1(input)
    util.CopyToClipboard(fmt.Sprintf("%v", ans))
    fmt.Println("Output:", ans)
  } else {
    ans := part2(input)
    util.CopyToClipboard(fmt.Sprintf("%v", ans))
    fmt.Println("Output:", ans)
  }
}

func part1(input string) int {
  walls, start, end, limit := parseInput(input)

  distance := generateDistFunc(walls)

  draw := func(path []directionalPos) {
    drawMap(walls, path, limit)
  }

  path := aStar(start, end, heuristic, distance, draw)
  // fmt.Println(start)
  // fmt.Println(path)
  // fmt.Println(len(path))

  drawMap(walls, path, limit)

  return calculateTrip(path)
}

func part2(input string) int {
  walls, start, end, _ := parseInput(input)

  distanceFunc := generateDistFunc(walls)
  
  cameFrom := djikstra(start,end, distanceFunc)
  possibleVisits := make(map[Pos]bool)
  possibleVisits[end] = true
  currentList := cameFrom[directionalPos{end, 0}]
  fmt.Println(currentList)
  for len(currentList) != 0 {
    nextList := make([]directionalPos,0)
    for _, dPos := range currentList {
      nextList = append(nextList, cameFrom[dPos]...)
      possibleVisits[dPos.Pos] = true
    }
    currentList = nextList
  
  }

  return len(possibleVisits)
}

func parseInput(input string) (walls map[Pos]bool, start directionalPos, end Pos, limit Pos) {
  walls = make(map[Pos]bool)
  for y, line := range strings.Split(input, "\n") {
    for x, char := range line {
      switch char {
      case '#':
        walls[Pos{x,y}] = true
      case 'S':
        start = directionalPos{Pos:Pos{x,y}, direction: 1}
      case 'E':
        end = Pos{x,y}
      }
    }
    limit = Pos{y+1, len(line)}
  }
  return walls, start, end, limit
}

type Pos = help.Pos

type directionalPos struct {
  Pos
  direction int
  // n = 0, e = 1, s = 2, w = 3
}

func (p directionalPos) OldNeighbors() []directionalPos {
  arr := make([]directionalPos, 0)
  arr = append(arr, directionalPos{p.Down(), 2})
  arr = append(arr, directionalPos{p.Left(), 3})
  arr = append(arr, directionalPos{p.Up(), 0})
  arr = append(arr, directionalPos{p.Right(), 1})
  return arr
}

func generateDistFunc(walls map[Pos]bool) (func (s directionalPos, e Pos) (cost int, success bool)) {
  distance := func (s directionalPos, e Pos) (cost int, success bool) {
    if walls[e] { return 0, false }
    cost = 1
    if ((s.direction % 2 == 0) && (e.Right() == s.Pos || e.Left() == s.Pos)) ||
       ((s.direction % 2 == 1) && (e.Up() == s.Pos || e.Down() == s.Pos)) {
      cost += 1000
    }
    if (s.direction == 0 && e.Up() == s.Pos)  || 
       (s.direction == 1 && e.Right() == s.Pos)  ||
       (s.direction == 2 && e.Down() == s.Pos)    ||
       (s.direction == 3 && e.Left() == s.Pos) {
      cost += 2000
    }
    return cost, true
  }

  return distance
}

func (p directionalPos) Neighbors() []directionalPos {
  arr := make([]directionalPos, 0)
  if p.direction != 0 { arr = append(arr, directionalPos{p.Down(), 2}) }
  if p.direction != 1 { arr = append(arr, directionalPos{p.Left(), 3}) }
  if p.direction != 2 { arr = append(arr, directionalPos{p.Up(), 0}) }
  if p.direction != 3 { arr = append(arr, directionalPos{p.Right(), 1}) }
  return arr
}

func heuristic(d directionalPos, goal Pos) (weight int) {
  deltaX := d.X - goal.X
  deltaY := d.Y - goal.Y
  
  switch d.direction {
  case 0:
    if deltaX != 0 { weight += 1000 }
    if deltaY <  0 { weight = 2000 }
  case 1:
    if deltaY != 0 { weight += 1000 }
    if deltaX >  0 { weight = 2000 }
  case 2:
    if deltaX != 0 { weight += 1000 }
    if deltaY >  0 { weight = 2000 }
  case 3:
    if deltaY != 0 { weight += 1000 }
    if deltaX <  0 { weight = 2000 }
  }

  if deltaX < 0 { deltaX *= -1 }
  if deltaY < 0 { deltaY *= -1 }
  weight += deltaX + deltaY
  return weight
}

func calculateTrip(path []directionalPos) int {
  sum := 0
  prev := path[0]

  for _, current := range path[1:] {
    if prev.direction != current.direction { sum += 1000 }
    sum += 1
    prev = current
  }

  return sum
}

func drawMap(walls map[Pos]bool, path []directionalPos, limit Pos) {
  pathMap := make(map[Pos]string)
  for key, _ := range walls {
    pathMap[key] = "#"
  }
  for _, pos := range path {
    switch pos.direction {
    case 0:
      pathMap[pos.Pos] = "^"
    case 1:
      pathMap[pos.Pos] = ">"
    case 2:
      pathMap[pos.Pos] = "v"
    case 3:
      pathMap[pos.Pos] = "<"
    }
  }
  for y := range limit.Y {
    lineString := ""
    for x := range limit.X {
      if val, ok := pathMap[Pos{x,y}]; ok {
        lineString += val
      } else {
        lineString += "."
      }
    }
    fmt.Println(lineString)
  }
}

func djikstra(start directionalPos, goal Pos, d distFunc) map[directionalPos][]directionalPos {

  cameFrom := make(map[directionalPos][]directionalPos)
  visited := make(map[directionalPos]bool)

  distance := make(map[directionalPos]int)
  distance[start] = 0

  unvisited := help.NewHeap[directionalPos](func (a,b directionalPos) bool {return distance[a] < distance[b] })
  unvisited.Push(start)

  for unvisited.Len() > 0 {
    current := unvisited.Pop()
    visited[current] = true

    for _, neighbor := range current.Neighbors() {
      if visited[neighbor] { continue }
      if neighbor.Pos == goal {
        neighbor.direction = 0
      }
      dist, success := d(current, neighbor.Pos)
      if !success {
        continue
      }
      tentativeScore := distance[current] + dist
      val, ok := distance[neighbor]
      if !ok {
        cameFrom[neighbor] = make([]directionalPos,0)
        cameFrom[neighbor] = append(cameFrom[neighbor], current)
        distance[neighbor] = tentativeScore
        unvisited.Push(neighbor)
      } else if tentativeScore == val {
        cameFrom[neighbor] = append(cameFrom[neighbor], current)
      } else if tentativeScore < val {
        cameFrom[neighbor] = make([]directionalPos,0)
        cameFrom[neighbor] = append(cameFrom[neighbor], current)
        distance[neighbor] = tentativeScore
      }
    }
    
  }

  return cameFrom
}
