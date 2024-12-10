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
  files, spaces := parseInput(input)
  done := false

  for i, file := range slices.Backward(files) {
    if done { break }
    new_blocks := make([]int, 0)
    removed_blocks:= make([]int, 0)
    for i, block := range slices.Backward(file.blocks) {
      if block < spaces[0] {
        new_blocks = slices.Concat(new_blocks, file.blocks[:i+1])
        done = true
        break
      }
      var new_block int
      new_block, spaces = spaces[0], spaces[1:]
      new_blocks = append(new_blocks, new_block)
      removed_blocks = append(removed_blocks, block)
    }

    file.blocks = new_blocks
    spaces = slices.Concat(spaces, removed_blocks)
    slices.Sort(spaces)

    files[i] = file
  }

  return getChecksum(files)
}

func part2(input string) int {
  files, space_blocks := parseInput(input)
  spaces := parseSpaces(space_blocks)

  for i, file := range slices.Backward(files) {
    for j, space := range spaces {
      if file.blocks[0] < space.start {
        break
      }
      if space.size >= file.size {
        new_blocks := make([]int, 0)

        for i := range file.size {
          new_blocks = append(new_blocks, space.start+i)
        }

        space.size -= file.size
        space.start += file.size
        spaces[j] = space

        file.blocks = new_blocks
        files[i] = file
        break
      }
    }
  }

  return getChecksum(files)
}

func parseInput(input string) (files []File, empties []int) {
  block := 0
  file_id := 0
  file := true
  files = make([]File, 0)
  empties = make([]int, 0)
  for _, char := range input {
    val := cast.ToInt(char)
    blocks := make([]int, 0)
      for _ = range val {
        blocks = append(blocks, block)
        block += 1
      }
    if file {
      files = append(files, File{file_id, val, blocks})
      file_id += 1
    } else {
      empties = slices.Concat(empties, blocks)
    }
    file = !file
  }
  return files, empties
}

type File struct {
  id int
  size int
  blocks []int
}

type Space struct {
  size int
  start int
}

func getChecksum(files []File) (sum int) {
  for _, file := range files {
    for _, block := range file.blocks {
      sum += block * file.id
    }
  }

  return sum
}

func parseSpaces(space_blocks []int) (spaces []Space) {
  spaces = make([]Space, 0)
  space := Space{1,space_blocks[0]}
  for _, block := range space_blocks[1:] {
    if block == space.start+space.size {
      space.size += 1
    } else {
      spaces = append(spaces, space)
      space = Space{1, block}
    }
  }
  spaces = append(spaces, space)

  return spaces
}
