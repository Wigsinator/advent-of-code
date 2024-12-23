package main

import (
  "testing"
)

var example1 = `AAAA
BBCD
BBCC
EEEC`

var example2 = `OOOOO
OXOXO
OOOOO
OXOXO
OOOOO`

var example = `RRRRIICCFF
RRRRIICCCF
VVRRRCCFFF
VVRCCCJFFF
VVVVCJJCFE
VVIVCCJJEE
VVIIICJJEE
MIIIIIJJEE
MIIISIJEEE
MMMISSJEEE`

func Test_part1(t *testing.T) {
  tests := []struct {
    name  string
    input string
    want  int
  }{
    {
      name:  "example 1",
      input: example1,
      want:  140,
    },
    {
      name:  "example 2",
      input: example2,
      want:  772,
    },
    {
      name:  "example 3",
      input: example,
      want:  1930,
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
      input: example,
      want:  1206,
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

func Test_bfs(t *testing.T) {
  parsed := parseInput(example1)
  visited := make(map[Pos]bool)

  _, r := bfs(Pos{0,0}, parsed, visited)
  want := Region{}
  want.area = 4
  want.perimiter = 10
  want.corners = 4
  if r != want {
    t.Errorf("bfs() = %v; want Region{4,10}", r)
  }
}
