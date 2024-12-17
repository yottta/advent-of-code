package _0_aoc

type Point struct {
	X, Y int
}

func (p *Point) AdjacentDown() Point {
	return Point{
		X: p.X,
		Y: p.Y + 1,
	}
}

func (p *Point) AdjacentUp() Point {
	return Point{
		X: p.X,
		Y: p.Y - 1,
	}
}

func (p *Point) AdjacentRight() Point {
	return Point{
		X: p.X + 1,
		Y: p.Y,
	}
}

func (p *Point) AdjacentLeft() Point {
	return Point{
		X: p.X - 1,
		Y: p.Y,
	}
}

func (p *Point) AllAdjacent() []Point {
	return []Point{
		p.AdjacentUp(),
		p.AdjacentRight(),
		p.AdjacentDown(),
		p.AdjacentLeft(),
	}
}

func (p *Point) InRange(minX, minY, maxX, maxY int) bool {
	if p.X < minX {
		return false
	}
	if p.Y < minY {
		return false
	}
	if p.X >= maxX {
		return false
	}
	if p.Y >= maxY {
		return false
	}
	return true
}
