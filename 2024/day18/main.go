package main

import (
  _ "embed"
  "flag"
  "fmt"
  "strings"

  "github.com/Wigsinator/advent-of-code/util"
  "github.com/Wigsinator/advent-of-code/cast"
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
    ans := part1(input, 70, 1024)
    util.CopyToClipboard(fmt.Sprintf("%v", ans))
    fmt.Println("Output:", ans)
  } else {
    ans := part2(input, 70)
    util.CopyToClipboard(fmt.Sprintf("%v", ans))
    fmt.Println("Output:", ans)
  }
}

func part1(input string, size int, drops int) int {
  parsed := parseInput(input)
  limit := Pos{size,size}
  corrupted := make(map[Pos]bool)
  for i := range drops {
    corrupted[parsed[i]] = true
  }
  heurFunc := genHeuristicFunc(limit)
  distFunc := genDistanceFunc(corrupted, limit)

  path := help.AStar(Pos{0,0}, limit, heurFunc, distFunc)

  return len(path)-1
}

func part2(input string, size int) string {
  parsed := parseInput(input)
  limit := Pos{size,size}
  hasPath := help.Memoize[int, bool](genHasPath(parsed, limit))

  left := 0
  right := len(parsed)-1
  var middle int

  for left < right {
    middle = (left + right) / 2
    check := isSolution(middle, hasPath)
    if check > 0 {
      left = middle + 1
    } else if check < 0 {
      right = middle - 1
    } else {
      break
    }
  }
  if left == right {
    middle = right
  }

  return fmt.Sprintf("%v,%v",parsed[middle].X, parsed[middle].Y)
}

func parseInput(input string) (ans []Pos) {
  ans = make([]Pos, 0)
  for _, line := range strings.Split(input, "\n") {
    vars := strings.Split(line, ",")
    ans = append(ans, Pos{cast.ToInt(vars[0]),cast.ToInt(vars[1])})
  }
  return ans
}

func genHeuristicFunc(limit Pos) (func(Pos) int) {
  heuristicFunc := func(a Pos) int {
    return (limit.Y-a.Y)+ (limit.X-a.X)
  }
  return heuristicFunc
}

// Because this is only being used in an A* system where we check neighbors, Distance is always 1 if the step is valid.
func genDistanceFunc(walls map[Pos]bool, limit Pos) (func(Pos,Pos) (int, bool)) {
  distFunc := func(a Pos, b Pos) (int, bool) {
    if walls[b] {
      return 0, false
    }
    if  b.X < 0 ||
        b.Y < 0 ||
        b.X > limit.X ||
        b.Y > limit.Y {
      return 0, false
    }
    return 1, true
  }
  return distFunc
}

func genHasPath(corruption []Pos, limit Pos) (func(int) bool) {
  heurFunc := genHeuristicFunc(limit)
  hasPath := func(middle int) bool {
    corrupted := make(map[Pos]bool)
    for i := range middle+1 {
      corrupted[corruption[i]] = true
    }
    distFunc := genDistanceFunc(corrupted, limit)
    path := help.AStar(Pos{0,0}, limit, heurFunc, distFunc)

    return (path != nil)
  }
  return hasPath
}

func isSolution(middle int, hasPath *help.Memoized[int,bool]) int {
  if hasPath.Call(middle) {
    return 1
  } else if hasPath.Call(middle-1){
    return 0
  } else {
    return -1
  }
}
