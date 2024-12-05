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

func part1(input string) (count int) {
  parsed := parseInput(input)
  for i, line := range parsed {
    for j, char := range line {
      if char == 'X' {
        if j >= 3 && parsed[i][j-1] == 'M' && parsed[i][j-2] == 'A' && parsed[i][j-3] == 'S' {
          count += 1
        }
        if j < len(parsed[i])-3 && parsed[i][j+1] == 'M' && parsed[i][j+2] == 'A' && parsed[i][j+3] == 'S' {
          count += 1
        }
        if i >= 3 && j >= 3 && parsed[i-1][j-1] == 'M' && parsed[i-2][j-2] == 'A' && parsed[i-3][j-3] == 'S' {
          count += 1
        }
        if i >= 3 && parsed[i-1][j] == 'M' && parsed[i-2][j] == 'A' && parsed[i-3][j] == 'S' {
          count += 1
        }
        if i >=3 && j < len(parsed[i])-3 && parsed[i-1][j+1] == 'M' && parsed[i-2][j+2] == 'A' && parsed[i-3][j+3] == 'S' {
          count += 1
        }
        if i < len(parsed)-3 && j >= 3 && parsed[i+1][j-1] == 'M' && parsed[i+2][j-2] == 'A' && parsed[i+3][j-3] == 'S' {
          count += 1
        }
        if i < len(parsed)-3 && parsed[i+1][j] == 'M' && parsed[i+2][j] == 'A' && parsed[i+3][j] == 'S' {
          count += 1
        }
        if i < len(parsed)-3 && j < len(parsed[i])-3 && parsed[i+1][j+1] == 'M' && parsed[i+2][j+2] == 'A' && parsed[i+3][j+3] == 'S' {
          count += 1
        }
      }
    }
  }

  return count
}

func part2(input string) (count int) {
  parsed := parseInput(input)
  for i, line := range parsed {
    if i == 0 || i+1 == len(parsed) {
      continue
    }
    for j, char := range line {
      if j == 0 || j+1 == len(line) {
        continue
      }
      if char == 'A' {
        if (parsed[i-1][j-1] == 'M' && parsed[i+1][j+1] == 'S') || (parsed[i-1][j-1] == 'S' && parsed[i+1][j+1] == 'M') {
          if (parsed[i-1][j+1] == 'M' && parsed[i+1][j-1] == 'S') || (parsed[i-1][j+1] == 'S' && parsed[i+1][j-1] == 'M') {
            count += 1
          }
        }
      }
    }
  }

  return count
}

func parseInput(input string) (ans []string) {
  for _, line := range strings.Split(input, "\n") {
    ans = append(ans, line)
  }
  return ans
}
