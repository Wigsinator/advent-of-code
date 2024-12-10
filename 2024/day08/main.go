package main

import (
  _ "embed"
  "flag"
  "fmt"
  "strings"

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
  chart, limit := parseInput(input)
  antinodes := make(map[Pos]bool)
  
  for _, antannae := range chart {
    for i, loc1 := range antannae[:len(antannae)-1] {
      for _, loc2 := range antannae[i+1:] {
        antis := findAntinodes(loc1, loc2)
        if antis[0].within(limit) {
          antinodes[antis[0]] = true
        }
        if antis[1].within(limit) {
          antinodes[antis[1]] = true
        }
      }
    }
  }

  return len(antinodes)
}

func part2(input string) int {
  chart, limit := parseInput(input)
  antinodes := make(map[Pos]bool)
  
  for _, antannae := range chart {
    for i, loc1 := range antannae[:len(antannae)-1] {
      for _, loc2 := range antannae[i+1:] {
        antis := findAntinodesPart2(loc1, loc2, limit)
        for _, anti := range antis {
          antinodes[anti] = true
        }
      }
    }
  }

  return len(antinodes)
}

func parseInput(input string) (ans map[rune][]Pos, limit Pos) {
  lines := strings.Split(input, "\n")
  limit = Pos{len(lines[0]),len(lines)}
  ans = make(map[rune][]Pos)
  for y, line := range lines {
    for x, char := range line {
      if char == '.' { continue }
      if _,ok := ans[char]; !ok {
        ans[char] = make([]Pos, 0)
      }
      ans[char] = append(ans[char], Pos{x,y})
    }
  }
  return ans, limit
}

type Pos struct {
  x,y int
}

func (p1 Pos) add(p2 Pos) Pos {
  return Pos{p1.x + p2.x, p1.y + p2.y}
}

func (p1 Pos) subtract(p2 Pos) Pos {
  return Pos{p1.x - p2.x, p1.y - p2.y}
}

func (p Pos) within(limit Pos) bool {
  return p.x >= 0 && p.y >= 0 && p.x < limit.x && p.y < limit.y
}

func findAntinodes(p1 Pos, p2 Pos) (antis [2]Pos) {
  antis[0] = Pos{(2 * p2.x) - p1.x, (2 * p2.y) - p1.y}
  antis[1] = Pos{(2 * p1.x) - p2.x, (2 * p1.y) - p2.y}
  return antis
}

func findAntinodesPart2(p1 Pos, p2 Pos, limit Pos) (antis []Pos) {
  antis = make([]Pos, 0)
  delta := p2.subtract(p1)

  current := p2
  for current.within(limit) {
    antis = append(antis, current)
    current = current.add(delta)
  }

  current = p1
  for current.within(limit) {
    antis = append(antis, current)
    current = current.subtract(delta)
  }

  return antis
}
