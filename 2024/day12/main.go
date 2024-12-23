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
    ans := part1(input)
    util.CopyToClipboard(fmt.Sprintf("%v", ans))
    fmt.Println("Output:", ans)
  } else {
    ans := part2(input)
    util.CopyToClipboard(fmt.Sprintf("%v", ans))
    fmt.Println("Output:", ans)
  }
}

func part1(input string) (total int) {
  parsed := parseInput(input)
  visited := make(map[Pos]bool)

  for p, _ := range parsed {
    if visited[p] { continue }

    var r Region

    visited, r = bfs(p, parsed, visited)
    // fmt.Println(parsed[p], r)
    total += r.area * r.perimiter
  }

  return total
}

func part2(input string) (total int) {
  parsed := parseInput(input)
  visited := make(map[Pos]bool)

  for p,_ := range parsed {
    if visited[p] {continue}

    var r Region

    visited, r = bfs(p,parsed,visited)
    total += r.area * r.corners
  }

  return total
}

func parseInput(input string) (ans map[Pos]rune) {
  ans = make(map[Pos]rune)
  for y, line := range strings.Split(input, "\n") {
    for x, char := range line {
      ans[Pos{x,y}] = char
    }
  }
  return ans
}

type Region struct {
  area, perimiter, corners int
}

func bfs (p Pos, plants map[Pos]rune, visited map[Pos]bool) (map[Pos]bool, Region) {
  queue := NewQueue()
  r := Region{}
  r_map := make(map[Pos]bool)
  queue.Push(p)
  visited[p] = true
  r_map[p] = true

  for queue.len() > 0 {
    // fmt.Println(queue, r)
    current := queue.Pop()
    r.area += 1
    r.corners += countCorners(current, plants)

    for _, nextPos := range current.Neighbors() {
      // fmt.Println(nextPos)
      if r_map[nextPos] { continue }
      if visited[nextPos] || plants[nextPos] != plants[current] {
        r.perimiter += 1
        continue
      }

      queue.Push(nextPos)
      visited[nextPos] = true
      r_map[nextPos] = true
    }
  }
  return visited, r
}

type Queue []Pos

func (self *Queue) Push(v Pos){
  *self = append(*self, v)
}

func (self *Queue) Pop() (val Pos){
  if len(*self) == 0 { panic("Can't pop empty queue") }
  h := *self
  val, *self = h[0], h[1:]
  return val
}

func NewQueue() *Queue {
  return &Queue{}
}

func (self *Queue) len() int {
  return len(*self)
}

func countCorners(p Pos, plants map[Pos]rune) (corners int) {
  surroundings := checkSurroundings(p, plants)
  neighbors := [4]bool{}
  for i := range 4 {
    neighbors[i] = surroundings[i]
  }

  switch neighbors {
  case [4]bool{false,false,false,false}:
    corners = 4
  case [4]bool{false,false,false,true}, [4]bool{false,false,true,false}, [4]bool{false,true,false,false}, [4]bool{true,false,false,false}:
    corners = 2
  case [4]bool{true,true,false,false}, [4]bool{true,false,true,false}, [4]bool{false,true,false,true}, [4]bool{false,false,true,true}:
    corners = 1
  }
  corners += checkInnerCorners(surroundings)

  return corners
}

func checkSurroundings(p Pos, plants map[Pos]rune) [8]bool {
  surroundings := [8]bool{}
  surroundings[0] = plants[p] == plants[p.Up()]
  surroundings[1] = plants[p] == plants[p.Left()]
  surroundings[2] = plants[p] == plants[p.Right()]
  surroundings[3] = plants[p] == plants[p.Down()]
  surroundings[4] = plants[p] == plants[p.Up().Left()]
  surroundings[5] = plants[p] == plants[p.Up().Right()]
  surroundings[6] = plants[p] == plants[p.Down().Left()]
  surroundings[7] = plants[p] == plants[p.Down().Right()]
  return surroundings
}

func checkInnerCorners(surr [8]bool) (val int) {
  if surr[0] && surr[1] && !surr[4] {val += 1}
  if surr[0] && surr[2] && !surr[5] {val += 1}
  if surr[3] && surr[1] && !surr[6] {val += 1}
  if surr[3] && surr[2] && !surr[7] {val += 1}
  return val
}
