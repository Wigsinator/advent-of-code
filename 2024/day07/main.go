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
  parsed := parseInput(input)
  
  for _,nums := range parsed {
    key, nums := nums[0], nums[1:]
    if valid_calibration(key, nums[1:], nums[0], false) {
      sum += key
    }
  }

  return sum
}

func part2(input string) (sum int) {
  parsed := parseInput(input)

  for _,nums := range parsed {
    key, nums := nums[0], nums[1:]
    if valid_calibration(key, nums[1:], nums[0], true) {
      sum += key
    }
  }

  return sum
}

func parseInput(input string) (ans [][]int) {
  ans = make([][]int, 0)
  for _, line := range strings.Split(input, "\n") {
    nums := make([]int, 0)
    vals := strings.Split(line, " ")
    key := cast.ToInt(vals[0][:len(vals[0])-1])
    nums = append(nums, key)
    for _, val := range vals[1:] {
      nums = append(nums, cast.ToInt(val))
    }
    ans = append(ans, nums)
  }
  return ans
}

func valid_calibration(key int, vals []int, current int, part2 bool) bool {
  if current > key { return false }
  if len(vals) == 0 {
    return key == current
  }

  var concatValid bool
  sumValid := valid_calibration(key, vals[1:], current +vals[0], part2)
  productValid := valid_calibration(key, vals[1:], current * vals[0], part2)
  if part2 {
    concatValue := cast.ToInt(cast.ToString(current) + cast.ToString(vals[0]))
    concatValid = valid_calibration(key, vals[1:], concatValue, part2)
  } else {
    concatValid = false
  }

  return sumValid || productValid || concatValid
}

