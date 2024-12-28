package main

import (
  "testing"
)

var exampleSmall = `########
#..O.O.#
##@.O..#
#...O..#
#.#.O..#
#...O..#
#......#
########

<^^>>>vv<v>>v<<`

var exampleLarge = `##########
#..O..O.O#
#......O.#
#.OO..O.O#
#..O@..O.#
#O#..O...#
#O..O..O.#
#.OO.O.OO#
#....O...#
##########

<vv>^<v^>v>^vv^v>v<>v^v<v<^vv<<<^><<><>>v<vvv<>^v^>^<<<><<v<<<v^vv^v>^
vvv<<^>^v^^><<>>><>^<<><^vv^^<>vvv<>><^^v>^>vv<>v<<<<v<^v>^<^^>>>^<v<v
><>vv>v^v^<>><>>>><^^>vv>v<^^^>>v^v^<^^>v^^>v^<^v>v<>>v^v^<v>v^^<^^vv<
<<v<^>>^^^^>>>v^<>vvv^><v<<<>^^^vv^<vvv>^>v<^^^^v<>^>vvvv><>>v^<<^^^^^
^><^><>>><>^^<<^^v>>><^<v>^<vv>>v>>>^v><>^v><<<<v>>v<v<v>vvv>^<><<>^><
^>><>^v<><^vvv<^^<><v<<<<<><^v<<<><<<^^<v<^^^><^>>^<v^><<<^>>^v<v^v<v^
>^>>^v>vv>^<<^v<>><<><<v<<v><>v<^vv<<<>^^v^>^^>>><<^v>>v^v><^^>>^<>vv^
<><^^>^^^<><vvvvv^v<v<<>^v<v>v<<^><<><<><<<^^<<<^<<>><<><^^^>^^<>^>v<>
^^>vv<^v^v<vv>^<><v<^v>^^^>>>^^vvv^>vvv<>>>^<^>>>>>^<<^v>^vvv<>^<><<v>
v^^>>><<^^<>>^v^<v^vv<>v^<<>^<^v^v><^<<<><<^<v><v<>vv>>v><v^<vv<>v^<<^`

var examplePart2 = `#######
#...#.#
#.....#
#..OO@#
#..O..#
#.....#
#######

<vv<<^^<<^^`

func Test_part1(t *testing.T) {
  tests := []struct {
    name  string
    input string
    want  int
  }{
    {
      name:  "exampleSmall",
      input: exampleSmall,
      want:  2028,
    },
    {
      name:   "exampleLarge",
      input:  exampleLarge,
      want:   10092,
    },
    // {
    //  name:  "actual",
    //  input: input,
    //  want:  0,
    // },
  }
  for _, tt := range tests {
    t.Run(tt.name, func(t *testing.T) {
      if got := part1(tt.input); got != tt.want {
        t.Errorf("part1() = %v, want %v", got, tt.want)
      }
    })
  }
}

func Test_part2(t *testing.T) {
  tests := []struct {
    name  string
    input string
    want  int
  }{
    {
      name:  "example",
      input: exampleLarge,
      want:  9021,
    },
    // {
    //  name:  "actual",
    //  input: input,
    //  want:  0,
    // },
  }
  for _, tt := range tests {
    t.Run(tt.name, func(t *testing.T) {
      if got := part2(tt.input); got != tt.want {
        t.Errorf("part2() = %v, want %v", got, tt.want)
      }
    })
  }
}
