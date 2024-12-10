package main

import (
  _ "embed"
  "flag"
  "fmt"
  "strings"

  "github.com/Wigsinator/advent-of-code/util"
  "github.com/Wigsinator/advent-of-code/cast"
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

func part1(input string) (sum int) {
  topo, heads := parseInput(input)
  
  for _, head := range heads {
    sum += getScore(head, topo)
  }

  return sum
}

func part2(input string) (sum int) {
  topo, heads := parseInput(input)
  ratings := make(map[Pos]int)

  for _, head := range heads {
    ratings = getRatingMap(head, topo, ratings)
    sum += ratings[head]
  }

  return sum
}

func parseInput(input string) (topo map[Pos]int, heads []Pos) {
  topo = make(map[Pos]int)
  heads = make([]Pos,0)
  for y, line := range strings.Split(input, "\n") {
    for x, char := range line {
      val := cast.ToInt(char)
      p := Pos{x,y}
      topo[p] = val
      if val == 0 {
        heads = append(heads, p)
      }
    }
  }
  return topo, heads
}

type Pos struct {
  x,y int
}

func getScore(p Pos, topo map[Pos] int) int {
  start := make(map[Pos]bool)
  start[p] = true
  peaks := getReachablePeaks(start, topo)
  if peaks == nil {
    return 0
  }
  return len(peaks)
}

func getRatingMap(p Pos, topo map[Pos]int, ratings map[Pos]int) (map[Pos]int) {
  if topo[p] == 9 { ratings[p] = 1 }
  if _, ok := ratings[p]; ok { return ratings }
  sum := 0

  for _, next := range p.nextPos(topo) {
    if _, ok := ratings[next]; !ok {
      ratings = getRatingMap(next, topo, ratings)
    }
    sum += ratings[next]
  }
  ratings[p] = sum

  return ratings
}

func getReachablePeaks(current map[Pos]bool, topo map[Pos]int) map[Pos]bool {
  if len(current) == 0 {
    return nil
  }
  next_level := make(map[Pos]bool)
  for key,_ := range current {
    if topo[key] == 9 {
      return current
    }
    for _, pos := range key.nextPos(topo) {
      next_level[pos] = true
    }
  }

  return getReachablePeaks(next_level, topo)
}

func (p Pos) nextPos(topo map[Pos]int) ([]Pos) {
  valid_next := make([]Pos, 0)
  if topo[p.up()] == topo[p] + 1 {
    valid_next = append(valid_next, p.up())
  }
  if topo[p.down()] == topo[p] + 1 {
    valid_next = append(valid_next, p.down())
  }
  if topo[p.left()] == topo[p] + 1 {
    valid_next = append(valid_next, p.left())
  }
  if topo[p.right()] == topo[p] + 1 {
    valid_next = append(valid_next, p.right())
  }
  return valid_next
}

func (p Pos) up() Pos {
  return Pos{p.x, p.y-1}
}
func (p Pos) down() Pos {
  return Pos{p.x, p.y+1}
}
func (p Pos) left() Pos {
  return Pos{p.x-1, p.y}
}
func (p Pos) right() Pos {
  return Pos{p.x+1, p.y}
}

