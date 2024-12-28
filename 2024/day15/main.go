package main

import (
  _ "embed"
  "flag"
  "fmt"
  "strings"

  "github.com/Wigsinator/advent-of-code/util"
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

func part1(input string) int {
  walls, boxes, robot, path, _ := parseInput(input)
  
  for _, direction := range path {
    boxes, robot = parseMovePart1(walls, boxes, robot, direction)
  }

  return sumBoxGPSPart1(boxes)
}

func part2(input string) int {
  input = doubleSize(input)
  walls, boxes, robot, path, limit := parseInput(input)
  bigBoxes := parseBoxes(boxes)

  for _, direction := range path {
    bigBoxes, robot = parseMovePart2(walls, bigBoxes, robot, direction)
  }
  
  visualizeMap(walls, bigBoxes, robot, limit)

  return sumBoxGPSPart2(bigBoxes)
}

func parseInput(input string) (walls map[Pos]bool, boxes map[Pos]bool, robot Pos, path string, limit Pos) {
  walls = make(map[Pos]bool)
  boxes = make(map[Pos]bool)

  inputSplit := strings.Split(input, "\n\n")
  path = strings.Join(strings.Split(inputSplit[1], "\n"), "")
  mapArray := strings.Split(inputSplit[0], "\n") 
  limit = Pos{len(mapArray[0]), len(mapArray)}

  for y, line := range mapArray{
    for x, char := range line {
      if char == '#' {
        walls[Pos{x,y}] = true
      }
      if char == 'O' || char == '[' {
        boxes[Pos{x,y}] = true
      }
      if char == '@' {
        robot = Pos{x,y}
      }
    }
  }

  return walls, boxes, robot, path, limit
}

func Move(p Pos, direction rune) (target Pos) {
  switch direction {
  case '^': 
    target = p.Up()
  case 'v':
    target = p.Down()
  case '<':
    target = p.Left()
  case '>':
    target = p.Right()
  }
  return target
}

func parseMovePart1(walls map[Pos]bool, boxes map[Pos]bool, robot Pos, direction rune) (map[Pos]bool, Pos) {
  target := Move(robot, direction)

  for boxes[target] {
    target = Move(target, direction)
  }
  if walls[target] {
    // Movement failed. Return the same values back.
    return boxes, robot
  }
  boxes[target] = true
  robot = Move(robot, direction)
  delete(boxes, robot)

  return boxes, robot
}

func sumBoxGPSPart1(boxes map[Pos]bool) (sum int) {
  for pos, val := range boxes {
    if val {
      sum += pos.Y * 100 + pos.X
    }
  }

  return sum
}

func doubleSize(input string) (output string) {
  inputSplit := strings.Split(input, "\n\n")
  for _, char := range inputSplit[0] {
    switch char {
    case '.':
      output += ".."
    case '#':
      output += "##"
    case '@':
      output += "@."
    case 'O':
      output += "[]"
    case '\n':
      output += "\n"
    }
  }
  output += "\n\n"
  output += inputSplit[1]
  return output
}

func parseBoxes(boxes map[Pos]bool) map[Pos]int {
  bigBoxes := make(map[Pos]int)

  for key, _ := range boxes {
    bigBoxes[key] = 1
    bigBoxes[key.Right()] = 2
  }

  return bigBoxes
}

func parseMovePart2(walls map[Pos]bool, boxes map[Pos]int, robot Pos, direction rune) (map[Pos]int, Pos) {
  if isValidMove(walls, boxes, robot, direction) {
    robot = Move(robot,direction)
    // If Boxes don't need to move, exit early
    if boxes[robot] == 0 { return boxes, robot }

    current := Move(robot, direction)
    prev := boxes[robot]
    if direction == '>' || direction == '<' {
      for boxes[current] != 0 {
        boxes[current], prev = prev, boxes[current]
        current = Move(current, direction)
      }
      boxes[current] = prev
    } else {
      boxesToMove := make(map[Pos]int)
      prevBoxes := make(map[Pos]int)
      boxesToMove[robot] = boxes[robot]
      if boxes[robot] == 1 {
        boxesToMove[robot.Right()] = 2
      }
      if boxes[robot] == 2 {
        boxesToMove[robot.Left()] = 1
      }

      for len(boxesToMove) > 0 {
        nextBoxesToMove := make(map[Pos]int)
        nextPrevBoxes := make(map[Pos]int)
        for pos, val := range boxesToMove {
          next := Move(pos, direction)
          if boxes[next] == 1 {
            nextBoxesToMove[next] = 1
            nextBoxesToMove[next.Right()] = 2
          }
          if boxes[next] == 2 {
            nextBoxesToMove[next] = 2
            nextBoxesToMove[next.Left()] = 1
          }
          boxes[next] = val
          nextPrevBoxes[next] = val
          if prevBoxes[pos] == 0 {
            delete(boxes, pos)
          }
        }
        prevBoxes = nextPrevBoxes
        boxesToMove = nextBoxesToMove
      }

    }
    delete(boxes, robot)
  }

  return boxes, robot
}

func isValidMove(walls map[Pos]bool, boxes map[Pos]int, robot Pos, direction rune) bool {
  target := Move(robot, direction)

  // Quick Checks for immediate Failure or Success
  if walls[target] { return false }
  if boxes[target] == 0 { return true }
  
  // Simple Check if moving horizontally
  if direction == '>' || direction == '<' {
    for boxes[target] != 0 {
      target = Move(target, direction)
    }
    return !walls[target]
  }

  // Oh dear god.
  targets := make(map[Pos]bool)
  targets[target] = true
  if boxes[target] == 1 {
    targets[target.Right()] = true
  }
  if boxes[target] == 2 {
    targets[target.Left()] = true
  }
  
  for len(targets) > 0 {
    next_targets := make(map[Pos]bool)
    for pos, _ := range targets {
      if walls[pos] {
        return false
      }
      next_pos := Move(pos, direction)
      if boxes[pos] == 1 {
        next_targets[next_pos] = true
        next_targets[next_pos.Right()] = true
      }
      if boxes[pos] == 2 {
        next_targets[next_pos] = true
        next_targets[next_pos.Left()] = true
      }
    }
    targets = next_targets
  }
  return true
}

func sumBoxGPSPart2(boxes map[Pos]int) (sum int) {
  for pos, val := range boxes {
    if val == 1 {
      sum += pos.Y * 100 + pos.X
    }
  }
  return sum
}

func visualizeMap(walls map[Pos]bool, boxes map[Pos]int, robot Pos, limit Pos) {
  for y := range limit.Y {
    lineString := ""
    for x := range limit.X {
      current := Pos{x,y}
      if walls[current] {
        lineString += "#"
      } else if boxes[current] == 1 {
        lineString += "["
      } else if boxes[current] == 2 {
        lineString += "]"
      } else if current == robot { 
        lineString += "@"
      } else {
        lineString += "."
      }
    }
    fmt.Println(lineString)
  }
}
