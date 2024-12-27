package main

import (
  _ "embed"
  "flag"
  "fmt"
  "strings"
  "regexp"

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

  for _, machine := range parsed {
    sum += findLowestCostPart2(machine)
  }

  return sum
}

func part2(input string) (sum int) {
  parsed := parseInput(input)

  for _, machine := range parsed {
    machine.prize.X += 10000000000000
    machine.prize.Y += 10000000000000
    sum += findLowestCostPart2(machine)
  }

  return sum
}

func parseInput(input string) (ans []Machine) {
  regex, err := regexp.Compile(`Button A: X\+(\d*), Y\+(\d*)[.\n]*Button B: X\+(\d*), Y\+(\d*)[.\n]*Prize: X=(\d*), Y=(\d*)`)
  if err != nil { panic("Regex Error") }
  for _, machineString := range strings.Split(input, "\n\n") {
    machine := Machine{}
    parsedMachine := regex.FindStringSubmatch(machineString)
    machine.buttonA.X = cast.ToInt(parsedMachine[1])
    machine.buttonA.Y = cast.ToInt(parsedMachine[2])
    machine.buttonB.X = cast.ToInt(parsedMachine[3])
    machine.buttonB.Y = cast.ToInt(parsedMachine[4])
    machine.prize.X = cast.ToInt(parsedMachine[5])
    machine.prize.Y = cast.ToInt(parsedMachine[6])
    ans = append(ans, machine)
  }
  return ans
}

type Machine struct {
  buttonA, buttonB, prize Pos
}

func findLowestCostPart1(machine Machine) (retval int) {
  for xScale := range 101 {
    for yScale := range 101 {
      if machine.prize == help.Add(machine.buttonA.Scale(xScale), machine.buttonB.Scale(yScale)) {
        if retval == 0 || retval < xScale * 3 + yScale {
          retval = xScale * 3 + yScale
        }
      }
    }
  }

  return retval
}

/*
ax + by = c
y = (c - ax)/b
y = (c/b) - (a/b)x

If there exists a pair of positive integers x,y that satisfy this, then we have a solution. If there are multiple, then we solve for the smallest value of 3x + y. 

The question is how to solve for integers which satisfy.

Alternatively, is there a way to *quickly* prove that there is no combination which can satisfy this equation.

By utilizing the fact that it has to line up both on the x and y axis, I'm able to solve it.
*/
func findLowestCostPart2(machine Machine) int {
  bPresses := float64((machine.prize.Y * machine.buttonA.X) - (machine.prize.X * machine.buttonA.Y))/float64((machine.buttonB.Y*machine.buttonA.X) - (machine.buttonB.X * machine.buttonA.Y))

  if !isInteger(bPresses) {
    return 0
  }

  aPresses := float64(machine.prize.X - int(bPresses) * machine.buttonB.X)/float64(machine.buttonA.X)

  if !isInteger(aPresses) {
    return 0
  }

  return int(aPresses) * 3 + int(bPresses)
}

func isInteger(val float64) bool {
  return val == float64(int(val))
}
