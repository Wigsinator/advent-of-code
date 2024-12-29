package main

import (
  _ "embed"
  "flag"
  "fmt"
  "strings"
  "slices"
  "math"

  "github.com/Wigsinator/advent-of-code/util"
  "github.com/Wigsinator/advent-of-code/cast"
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

func part1(input string) string {
  prog := parseInput(input)

  prog.run()

  return prog.outputString()
}

func part2(input string) (retval int) {
  prog := parseInput(input)
  startVal := int(math.Pow(8,float64(len(prog.instructions)-1)))
  scoreMap := make(map[int]int)
  getScore := genGetScore(prog, scoreMap)
  nextFunc := generateNextValuesFunction(prog, getScore)
  validationFunc := generateValidationFunc(prog, getScore)
  current := startVal
  possibilites := make([]int, 0)
  for _ = range 7 {
    values := help.Dfs[int](current, nextFunc, validationFunc)
    possibilites = append(possibilites, values...)
    current += startVal
  }

  if len(possibilites) == 0 {
    retval = -1
  } else {
    retval = possibilites[0]
    for _, item := range possibilites[1:] {
      if item < retval {
        retval = item
      }
    }
  }

  return retval
}

func parseInput(input string) (ans *Program) {
  instructions := make([]int, 0)

  var regA, regB, regC int

  for i, line := range strings.Split(input, "\n") {
    lineSlice := strings.Split(line, " ")
    lastEl := lineSlice[len(lineSlice)-1]
    switch i {
    case 0:
      regA = cast.ToInt(lastEl)
    case 1:
      regB = cast.ToInt(lastEl)
    case 2:
      regC = cast.ToInt(lastEl)
    case 4:
      for _, instruction := range strings.Split(lastEl, ",") {
        instructions = append(instructions, cast.ToInt(instruction))
      }
    }
  }

  ans = newProgram(instructions)
  ans.setRegA(regA)
  ans.setRegB(regB)
  ans.setRegC(regC)
  return ans
}

func pow2(power int) int {
  return int(math.Pow(2, float64(power)))
}

type Program struct {
  pointer int
  regA, regB, regC int
  instructions []int
  output []int
  jumped bool
}

func newProgram(instructions []int) *Program {
  return &Program{pointer:0, regA:0, regB:0, regC:0, instructions: instructions, output: make([]int,0), jumped: false}
}

func (p *Program) setRegA(val int)  { p.regA = val }
func (p *Program) setRegB(val int)  { p.regB = val }
func (p *Program) setRegC(val int)  { p.regC = val }

func (p *Program) reset(A int) { 
  p.setRegA(A)
  p.setRegB(0)
  p.setRegC(0)
  p.pointer = 0
  p.output = make([]int,0)
}

// Returns -1 if Halt Required
func (p *Program) readOpcode() int {
  if p.pointer < len(p.instructions) {
    return p.instructions[p.pointer]
  } else {
    return -1
  }
}

func (p *Program) readOperand(combo bool) (output int) {
  literalOperand := p.instructions[p.pointer+1]
  if !combo { return literalOperand }
  switch literalOperand {
  case 0,1,2,3:
    output = literalOperand
  case 4:
    output = p.regA
  case 5:
    output = p.regB
  case 6:
    output = p.regC
  case 7:
    panic("Illegal combo Operand 7")
  }
  return output
}

func (p *Program) increment() {
  if p.jumped {
    p.jumped = false
  } else {
    p.pointer += 2
  }
}

// If validate == true, Exits early if output contains discrepancy with instructions
func (p *Program) run() {
  loop: for true {
    // p.printOut()
    switch p.readOpcode() {
    case -1:
      break loop
    case 0:
      p.adv()
    case 1:
      p.bxl()
    case 2:
      p.bst()
    case 3:
      p.jnz()
    case 4:
      p.bxc()
    case 5:
      p.out()
    case 6:
      p.bdv()
    case 7:
      p.cdv()
    }
    p.increment()
  }
}

// Helper function for adv, bdv, and cdv
func (p *Program) getDivValue() int {
  operand := p.readOperand(true)
  divisor := pow2(operand)
  return p.regA / divisor
}

// Opcode 0. Combo Operand
func (p *Program) adv() {
  p.setRegA(p.getDivValue())
}
// Opcode 6. Combo Operand
func (p *Program) bdv() {
  p.setRegB(p.getDivValue())
}
// Opcode 7. Combo Operand
func (p *Program) cdv() {
  p.setRegC(p.getDivValue())
}

// Opcode 1. Literal operand
func (p *Program) bxl() {
  operand := p.readOperand(false)
  value := p.regB ^ operand
  p.setRegB(value)
}

// Opcode 2. Combo Operand
func (p *Program) bst() {
  operand := p.readOperand(true)
  value := operand % 8
  p.setRegB(value)
}

// Opcode 3. Literal Operand
func (p *Program) jnz() {
  if p.regA != 0 {
    operand := p.readOperand(false)
    p.pointer = operand
    p.jumped = true
  }
}

// Opcode 4. None Operand
func (p *Program) bxc() {
  value := p.regB ^ p.regC
  p.setRegB(value)
}

// Opcode 5. Combo Operand
func (p *Program) out() {
  operand := p.readOperand(true)
  outval := operand % 8
  p.output = append(p.output, outval)
}

func (p *Program) outputString() string {
  return strings.Trim(strings.Join(strings.Split(fmt.Sprint(p.output), " "), ","), "[]")
}

func (p *Program) printOut() {
  fmt.Println("Register A:", p.regA)
  fmt.Println("Register B:", p.regB)
  fmt.Println("Register C:", p.regC)
  fmt.Println(strings.Repeat("  ",p.pointer), "v")
  fmt.Println(p.instructions)
  fmt.Println(p.output)
  fmt.Println("^ Current Output")
}

func (p *Program) validEndings() (count int) {
  if len(p.output) != len(p.instructions) {return 0}
  for i, val := range slices.Backward(p.output) {
    if val != p.instructions[i] {
      break
    }
    count += 1
  }
  return count
}

func generateNextValuesFunction(p *Program, getScore func(int) int) (func(int) []int) {
  nextValuesFunc := func(aVal int) []int {
    validEnds := getScore(aVal)
    maxPowToAlter := countOctalTrailingZeroes(aVal)-1
    fmt.Println("NextValues for", aVal, "with score of", validEnds)
    arr := make([]int,0)
    powToAlter := 0
    for powToAlter <= maxPowToAlter {
      increment := int(math.Pow(8,float64(powToAlter)))
      fmt.Println("Adding increments of", increment, "or 8^",powToAlter)
      for i := range 8 {
        possibleAVal := aVal + i * increment
        possibleScore := getScore(possibleAVal)
        fmt.Println("Verifying", possibleAVal, "with score of", possibleScore)
        if possibleScore + powToAlter >= len(p.instructions) {
          arr = append(arr, possibleAVal)
        }
      }
      powToAlter += 1
    }
    return arr
  }
  return nextValuesFunc
}

func generateValidationFunc(p *Program, getScore func(int) int) (func(int) bool) {
  validationFunc := func(aVal int) bool {
    return len(p.instructions) == getScore(aVal)
  }
  return validationFunc
}


func genGetScore(p *Program, scoreMap map[int]int) (func(int) int) {
  getScore := func(aVal int) int {
    if _, ok := scoreMap[aVal]; !ok {
      p.reset(aVal)
      p.run()
      fmt.Printf("Starting AVal of %v, Octal: %o\n", aVal, aVal)
      p.printOut()
      scoreMap[aVal] = p.validEndings()
    }
    return scoreMap[aVal]
  }
  return getScore
}

func countOctalTrailingZeroes(n int) int {
  count := 0
  for p := 8; n%p == 0; p*= 8 {
    count += 1
  }
  return count
}
