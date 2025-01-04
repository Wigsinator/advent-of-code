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

type Pos = help.Pos

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
    ans := part1(input, 100)
    util.CopyToClipboard(fmt.Sprintf("%v", ans))
    fmt.Println("Output:", ans)
  } else {
    ans := part2(input, 100)
    util.CopyToClipboard(fmt.Sprintf("%v", ans))
    fmt.Println("Output:", ans)
  }
}

func part1(input string, minimum int) int {
  walls, start, end, limit := parseInput(input)
  distFunc := genDistFunc(walls, limit)

  distMap, cameFrom := help.DjikstraPos(start, end, distFunc)

  current := end
  var next bool
  cheatCounts := make(map[int]int)

  for true {
    for _, neighbor := range current.Neighbors() {
      if !walls[neighbor] || !neighbor.Within(limit) { continue }
      for _, nNeighbor := range neighbor.Neighbors() {
        if _, ok := distMap[nNeighbor]; !ok { continue }
        cheatDist := distMap[current] - distMap[nNeighbor]-2 
        if cheatDist > 0 {
          // fmt.Println("Adding Cheat Distance:", cheatDist, "| From",current, "to", nNeighbor)
          cheatCounts[cheatDist] += 1
        }
      }
    }
    
    current, next = cameFrom[current]
    if !next { break }
  }

  count := 0
  for key, val := range cheatCounts {
    if key >= minimum {
      count += val
    }
  }

  // fmt.Println(cheatCounts)

  return count
}

func part2(input string, minimum int) int {
  walls, start, end, limit := parseInput(input)
  distFunc := genDistFunc(walls, limit)

  distMap, cameFrom := help.DjikstraPos(start, end, distFunc)

  current := end
  var next bool
  cheatCounts := make(map[int]int)
  cheatDists := make(map[cheat]int)

  for true {
    toCheck := make([]Pos,0)
    for dX := range 21 {
      for dY := range 21 - dX {
        toCheck = append(toCheck, Pos{current.X-dX, current.Y-dY})
        toCheck = append(toCheck, Pos{current.X+dX, current.Y-dY})
        toCheck = append(toCheck, Pos{current.X+dX, current.Y+dY})
        toCheck = append(toCheck, Pos{current.X-dX, current.Y+dY})
      }
    }

    for _, newPos := range toCheck {
      if _, ok := distMap[newPos]; !ok { continue }
      if _, ok := cheatDists[cheat{newPos,current}]; ok { continue }
      cheatDist := distMap[current] - distMap[newPos] - help.TaxiDist(current,newPos)
      if cheatDist >= 50 {
        cheatDists[cheat{newPos,current}] = cheatDist
        cheatCounts[cheatDist] += 1
      }  
    }
    
    current, next = cameFrom[current]
    if !next { break }
  }

  count := 0
  for key, val := range cheatCounts {
    if key >= minimum {
      count += val
    }
  }

  return count
}

func parseInput(input string) (walls map[Pos]bool, start Pos, end Pos, limit Pos) {
  lines := strings.Split(input, "\n")
  limit = Pos{len(lines), len(lines[0])}
  walls = make(map[Pos]bool)

  for y, line := range lines {
    for x, char := range line {
      switch char{
      case '#':
        walls[Pos{x,y}] = true
      case 'S':
        start = Pos{x,y}
      case 'E':
        end = Pos{x,y}
      }
    }
  }
  return walls, start, end, limit
}

func genDistFunc(walls map[Pos]bool, limit Pos) func(Pos,Pos) int {
  return func(a Pos, b Pos) int {
    if  walls[b] ||
        !b.Within(limit) { 
      return -1 
    } else {
      return 1
    }
  }
}

type cheat struct {
  start Pos
  end Pos
}
