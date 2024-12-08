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

func part1(input string) (sum int) {
  rules, updates := parseInput(input)
  rules_cmp := func(a,b int) int {
    return rules[[2]int{a,b}]
  }
  for _, update := range updates {
    if slices.IsSortedFunc(update, rules_cmp) {
      sum += update[(len(update)-1)/2]
    }
  }

  return sum
}

func part2(input string) (sum int) {
  rules, updates := parseInput(input)
  rules_cmp := func(a,b int) int {
    return rules[[2]int{a,b}]
  }
  for _, update := range updates {
    if !slices.IsSortedFunc(update, rules_cmp) {
      slices.SortFunc(update, rules_cmp)
      sum += update[(len(update)-1)/2]
    }
  }

  return sum
}

func parseInput(input string) (rules map[[2]int]int, updates [][]int) {
  inputs := strings.Split(input, "\n\n")
  rules = make(map[[2]int]int)
  for _, line := range strings.Split(inputs[0], "\n") {
    vals := strings.Split(line, "|")
    val_array := [2]int{cast.ToInt(vals[0]),cast.ToInt(vals[1])}
    rules[[2]int{val_array[1],val_array[0]}] = 1
    rules[val_array] = -1
  }
  for _, line := range strings.Split(inputs[1], "\n") {
    update := make([]int, 0)
    for _, val := range strings.Split(line, ",") {
      update = append(update, cast.ToInt(val))
    }
    updates = append(updates, update)
  }
  return rules, updates
}

/* DEPRECATED
func updateIsValid(rules map[int][]int, update []int) bool {
  for i,val := range update[:len(update)-1] {
    if has_intersection(rules[val], update[i+1:]) {
      return false
    }
  }
  return true
}

func has_intersection(a []int, b []int) bool {
  hash := make(map[int]bool)

  for _,v := range a {
    hash[v] = true
  }

  for _,v := range b {
    if hash[v] {
      return true
    }
  }

  return false
}
*/

