package main

import (
  _ "embed"
  "flag"
  "fmt"
  "strings"
  "regexp"

  "github.com/Wigsinator/advent-of-code/util"
  "github.com/Wigsinator/advent-of-code/help"
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
  towels, designs := parseInput(input)
  pattern := fmt.Sprintf("^(?:%v)+$", strings.Join(towels,"|"))
  regex, err := regexp.Compile(pattern)
  if err != nil { panic("Regex Error") }

  for _, design := range designs {
    if regex.MatchString(design) {
      count += 1
    }
  }

  return count
}

func part2(input string) (count int) {
  towels, designs := parseInput(input)
  pattern := fmt.Sprintf("^(?:%v)+$", strings.Join(towels,"|"))
  regex, err := regexp.Compile(pattern)
  if err != nil { panic("Regex Error") }
  memo := make(map[string]int)

  for _, design := range designs {
    if regex.MatchString(design) {
      // count += countWaysBFS(towels, design)
      count += countWaysRecurse(design, towels, memo)
    }
  }

  return count
}

func parseInput(input string) (towels []string, designs []string) {
  parts := strings.Split(input, "\n\n") 
  towels = strings.Split(parts[0], ", ")
  designs = strings.Split(parts[1],"\n")
  return towels, designs
}

func countWaysBFS(towels []string, design string) int {
  fmt.Println("Finding methods for", design)
  q := help.NewQueue[string]()
  parents := make(map[string][]string)
  cache := make(map[string]int)
  cache[""] = 1
  for _, towel := range addTowels("",towels,design) {
    q.Push(towel)
    parents[towel] = make([]string,0)
    parents[towel] = append(parents[towel], "")
  }

  for !q.IsEmpty() {
    val := q.Pop()
    // fmt.Println("Testing",val, "| Number of ways:", parents[val])
    if val == design {
      continue
    }
    for _, w := range addTowels(val, towels, design) {
      if _, ok := parents[w]; !ok {
        q.Push(w)
        parents[w] = make([]string,0)
      }
      parents[w] = append(parents[w], val)
    }  
  }

  return countParents(parents, design, cache)
}

func addTowels(val string, towels []string, design string) []string {
  retval := make([]string,0)
  for _, towel := range towels {
    newString := val + towel
    if strings.HasPrefix(design, newString) {
      retval = append(retval, newString)
    }
  }

  return retval
}

func countParents(parents map[string][]string, val string, cache map[string]int) (retval int) {
  if num, ok := cache[val]; ok {
    return num
  }
  
  for _, parent := range parents[val] {
    retval += countParents(parents, parent, cache)
  }

  cache[val] = retval
  return retval
}

func countWaysRecurse(design string, towels []string, memo map[string]int) (result int) {
  if design == "" { return 1 }
  if num, ok := memo[design]; ok { return num }
  
  for _, towel := range towels {
    if newDesign, ok := strings.CutPrefix(design, towel); ok {
      result += countWaysRecurse(newDesign, towels, memo)
    }
  }
  memo[design] = result
  return result
}
