package help

// Helper functions for Managing 2d Positions

type Pos struct {
  X,Y int
}

func (p Pos) Up() Pos {
  return Pos{p.X, p.Y-1}
}
func (p Pos) Down() Pos {
  return Pos{p.X, p.Y+1}
}
func (p Pos) Left() Pos {
  return Pos{p.X-1, p.Y}
}
func (p Pos) Right() Pos {
  return Pos{p.X+1, p.Y}
}

func (p Pos) Neighbors() []Pos {
  return []Pos{p.Up(), p.Down(), p.Left(), p.Right()}
}
