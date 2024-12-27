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

func (p Pos) Scale(factor int) Pos {
  return Pos{p.X * factor, p.Y * factor}
}

func Add(p1 Pos, p2 Pos) Pos {
  return Pos{p1.X + p2.X, p1.Y + p2.Y}
}

