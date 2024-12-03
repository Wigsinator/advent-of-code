package main

import (
  _ "embed"
  "flag"
  "fmt"
  "strings"
  "regexp"

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

func part1(input string) (ans int) {
  parsed := parseInput1(input)
  for _, mult := range parsed {
    ans += mult[0] * mult[1]
  }
  return ans
}

func part2(input string) (ans int) {
  parsed := parseInput2(input)
  for _, mult := range parsed {
    ans += mult[0] * mult[1]
  }
  return ans
}

func parseInput1(input string) (ans [][]int) {
  re := regexp.MustCompile(`mul\((\d{1,3}),(\d{1,3})\)`)
  matches := re.FindAllStringSubmatch(input, -1)
  for _, match := range matches {
    vals := []int{cast.ToInt(match[1]), cast.ToInt(match[2])}
    ans = append(ans, vals)
  }
  return ans
}

func parseInput2(input string) (ans [][]int) {
  re := regexp.MustCompile(`mul\((\d{1,3}),(\d{1,3})\)|do\(\)|don't\(\)`)
  matches := re.FindAllStringSubmatch(input, -1)
  active := true
  for _, match := range matches {
    if match[0] == "don't()" {
      active = false
    } else if match[0] == "do()" {
      active = true
    } else if active {
      vals := []int{cast.ToInt(match[1]), cast.ToInt(match[2])}
      ans = append(ans, vals)
    }
  }
  return ans
}
