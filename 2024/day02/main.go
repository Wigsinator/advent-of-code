package main

import (
  _ "embed"
  "flag"
  "fmt"
  "strings"
  "slices"

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

func part1(input string) (count int) {
  parsed := parseInput(input)
  for _, report := range parsed {
    if isSafe(report) {
      count += 1
    }
  }
  return count
}

func part2(input string) (count int) {
  parsed := parseInput(input)
  for _, report := range parsed {
    if isSafeWithDampener(report) {
      count += 1
    }
  }
  return count
}

func parseInput(input string) (ans [][]int) {
  for _, line := range strings.Split(input, "\n") {
    lineNums := make([]int, 0)
    for _, num := range strings.Split(line, " ") {
      lineNums = append(lineNums, cast.ToInt(num))
    }
    ans = append(ans, lineNums)
  }
  return ans
}

func isSafe(report []int) (bool) {
  asc := report[0] < report[1]
  for index, _ := range report {
    if index + 1 == len(report) {
      continue
    }

    diff := 0
    if asc {
      diff = report[index + 1] - report[index]
    } else {
      diff = report[index] - report[index + 1]
    }

    if (diff > 3 || diff < 1) {
      return false
    }
  }
  return true
}

func isSafeWithDampener(report []int) (bool) {
  if isSafe(report) {
    return true
  }
  for index, _ := range report {
    if isSafe(slices.Concat(report[:index],report[index+1:])){
      return true
    }
  }
  return false
}
