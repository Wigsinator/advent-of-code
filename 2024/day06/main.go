package main

import (
  _ "embed"
  "flag"
  "fmt"
  "strings"
  "maps"

  "github.com/Wigsinator/advent-of-code/util"
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
  obstacles, guard, limit := parseInput(input)
  visited := make(map[Pos]bool)
  
  for guard.inLimit(limit) {
    visited[guard.pos] = true
    guard = guard.move(obstacles)
  }

  return len(visited)
}

func part2(input string) (count int) {
  obstacles, guard, limit := parseInput(input)
  causes_loops := make(map[Pos]bool)
  
  for guard.inLimit(limit) {
    new_obs := guard.move(obstacles).pos
    if _, ok := causes_loops[new_obs]; !ok {
      causes_loops[new_obs] = loopAhead(guard, maps.Clone(obstacles), limit, new_obs)
    }
    guard = guard.move(obstacles)
  }

  for _,val := range causes_loops {
    if val { count += 1 }
  }

  return count
}

func parseInput(input string) (obstacles map[Pos]bool, guard Guard, limit Pos) {
  lines := strings.Split(input, "\n")
  obstacles = make(map[Pos]bool)
  limit = Pos{len(lines),len(lines[0])}
  for y, line := range lines {
    for x, val := range line {
      if val == '#' {
        obstacles[Pos{x,y}] = true
      }
      if val == '^' {
        guard = Guard{Pos{x,y}, [2]int{0,-1}}
      }
    }
  }
  return obstacles, guard, limit
}

type Pos struct {
  x,y int
}

type Guard struct {
  pos Pos
  direction [2]int
}

func (g Guard) rotate() Guard {
  g.direction = [2]int{-g.direction[1], g.direction[0]}
  return g
}

func (g Guard) move(obstacles map[Pos]bool) Guard {
  next_pos := Pos{g.pos.x + g.direction[0], g.pos.y + g.direction[1]}
  for obstacles[next_pos] {
    g = g.rotate()
    next_pos = Pos{g.pos.x + g.direction[0], g.pos.y + g.direction[1]}
  }
  g.pos = next_pos
  return g
}

func (g Guard) inLimit(limit Pos) bool { return g.pos.x >= 0 && g.pos.x < limit.x && g.pos.y >= 0 && g.pos.y < limit.y }

func loopAhead(guard Guard, obstacles map[Pos]bool, limit Pos, new_obs Pos) bool {
  obstacles[new_obs] = true
  guardstates := make(map[Guard]bool)

  for guard.inLimit(limit) && !guardstates[guard] {
    guardstates[guard] = true
    guard = guard.move(obstacles)
  }
  return guardstates[guard]
}
