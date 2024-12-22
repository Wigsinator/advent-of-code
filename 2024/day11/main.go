package main

import (
  _ "embed"
  "flag"
  "fmt"
  "strings"
  "math"

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
  parsed := parseInput(input)
  for range 25 {
    parsed = iterate_rules(parsed)
  }
  for _, val := range parsed {
    sum += val
  }

  return sum
}

func part2(input string) (sum int) {
  parsed := parseInput(input)
  for range 75 {
    parsed = iterate_rules(parsed)
  }
  for _, val := range parsed {
    sum += val
  }
  
  return sum
}

func parseInput(input string) (ans map[int]int) {
  ans = make(map[int]int)
  for _, num := range strings.Split(input, " ") {
    ans[cast.ToInt(num)] += 1
  }
  return ans
}

func iterate_rules(input map[int]int) (ans map[int]int) {
  ans = make(map[int]int)
  for key, val := range input {
    if key == 0 {
      ans[1] += val
    } else if length := int(math.Log10(float64(key))) + 1; length % 2 == 0 {
      ans[key/int(math.Pow10(length/2))] += val
      ans[key%int(math.Pow10(length/2))] += val
    } else {
      ans[key*2024] += val
    }
  }
  return ans
}
