package main

import (
  "testing"
)

var example1 = `Register A: 729
Register B: 0
Register C: 0

Program: 0,1,5,4,3,0`

var example2 = `Register A: 2024
Register B: 0
Register C: 0

Program: 0,3,5,4,3,0`

func Test_part1(t *testing.T) {
  tests := []struct {
    name  string
    input string
    want  string
  }{
    {
      name:  "example",
      input: example1,
      want:  "4,6,3,5,6,3,5,2,1,0",
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
      input: example2,
      want:  117440,
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
