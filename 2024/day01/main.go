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

func part1(input string) int {
	parsed1,parsed2 := parseInput(input)
	slices.Sort(parsed1)
  slices.Sort(parsed2)
  sum := 0
  for index, _ := range parsed1 {
    if (parsed1[index] > parsed2[index]) {
      sum += parsed1[index] - parsed2[index]
    } else {
      sum += parsed2[index] - parsed1[index]
    }
  }
	return sum
}

func part2(input string) int {
  parsed1, parsed2 := parseInput(input)
  counts := get_counts(parsed2)
  sum := 0
  for _, val := range parsed1 {
    sum += val * counts[val]
  }
  return sum
}

func parseInput(input string) (list1 []int, list2 []int) {
	for _, line := range strings.Split(input, "\n") {
    vals := strings.Split(line, "   ")
		list1 = append(list1, cast.ToInt(vals[0]))
    list2 = append(list2, cast.ToInt(vals[1]))
	}
	return list1, list2
}

func get_counts(slice []int) map[int]int {
  counts := make(map[int]int)
  for _, val := range slice {
    counts[val] += 1
  }
  return counts
}
