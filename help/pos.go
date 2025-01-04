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

func (p Pos) Within(limit Pos) bool {
  retval := p.X < limit.X ||
            p.Y < limit.Y ||
            p.X >= 0 ||
            p.Y >= 0
  return retval
}

func TaxiDist(p1 Pos, p2 Pos) int {
  dY := p2.Y - p1.Y
  if dY < 0 { dY *= -1 }
  dX := p2.X - p1.X
  if dX < 0 { dX *= -1 }
  return dY + dX
}
