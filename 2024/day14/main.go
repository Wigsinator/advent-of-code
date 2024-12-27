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
    ans := part1(input, 101, 103)
    util.CopyToClipboard(fmt.Sprintf("%v", ans))
    fmt.Println("Output:", ans)
  } else {
    ans := part2(input, 101, 103)
    util.CopyToClipboard(fmt.Sprintf("%v", ans))
    fmt.Println("Output:", ans)
  }
}

func part1(input string, width int, height int) int {
  parsed := parseInput(input)

  for i, robot := range parsed {
    parsed[i] = move(robot, 100, width, height)
  }

  return calculateSafetyScore(parsed, width, height)
}

func part2(input string, width int, height int) int {
  parsed := parseInput(input)
  timeElapsed := 0

  for !treeAppeared(parsed) {
    for i, robot := range parsed {
      parsed[i] = move(robot, 1, width, height)
    }
    timeElapsed += 1
  }
  visualizeGrid(parsed, width, height)

  return timeElapsed
}

func parseInput(input string) (ans []Robot) {
  regex, err := regexp.Compile(`p=(\d*),(\d*) v=(-?\d*),(-?\d*)`)
  if err != nil { panic("Regex Error") }
  for _, robotString := range strings.Split(input, "\n") {
    robotMatch := regex.FindStringSubmatch(robotString)
    robot := Robot{}
    robot.position.X = cast.ToInt(robotMatch[1])
    robot.position.Y = cast.ToInt(robotMatch[2])
    robot.velocity.X = cast.ToInt(robotMatch[3])
    robot.velocity.Y = cast.ToInt(robotMatch[4])
    ans = append(ans, robot)
  }
  return ans
}

type Robot struct {
  position, velocity Pos
}

func move(robot Robot, repetitions int, width int, height int) Robot {
  robot.position = help.Add(robot.position, robot.velocity.Scale(repetitions))
  robot.position.X = (robot.position.X % width + width) % width
  robot.position.Y = (robot.position.Y % height + height) % height

  return robot
}

func calculateSafetyScore(robots []Robot, width int, height int) int {
  q1,q2,q3,q4 := 0,0,0,0
  for _, robot := range robots {
    if robot.position.X < width / 2 {
      if robot.position.Y < height / 2 {
        q1 += 1
      }
      if robot.position.Y > height / 2 {
        q2 += 1
      }
    }
    if robot.position.X > width / 2 {
      if robot.position.Y < height / 2 {
        q3 += 1
      }
      if robot.position.Y > height / 2 {
        q4 += 1
      }
    }
  }

  return q1 * q2 * q3 * q4
}

func treeAppeared(robots []Robot) bool {
  locations := make(map[Pos]bool)
  for _, robot := range robots {
    locations[robot.position] = true
  }
  oneLine := false
  for _, robot := range robots {
    inline := 1
    current := robot.position
    for locations[current.Right()] {
      inline += 1
      current = current.Right()
    }
    if inline > 7 {
      if oneLine {
        return true
      } else {
        oneLine = true
      }
    }
  }

  return false
}

func visualizeGrid(robots []Robot, width int, height int) {
  locations := make(map[Pos]bool)
  for _, robot := range robots {
    locations[robot.position] = true
  }

  for y := range height {
    rowString := ""
    for x := range width {
      if locations[Pos{x,y}] {
        rowString += "■"
      } else {
        rowString += "☐"
      }
    }
    fmt.Println(rowString)
  }
}
