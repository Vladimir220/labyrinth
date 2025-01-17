package labyrinth

type Point struct {
	X, Y int
}

func (p Point) InAria(p1, p2 Point) bool {
	var minX, maxX, minY, maxY int

	if p1.X < p2.X {
		minX = p1.X
		maxX = p2.X
	} else {
		maxX = p1.X
		minX = p2.X
	}
	if p1.Y < p2.Y {
		minY = p1.Y
		maxY = p2.Y
	} else {
		maxY = p1.Y
		minY = p2.Y
	}

	if p.X < minX || p.X > maxX || p.Y < minY || p.Y > maxY {
		return false
	}

	return true
}

func (p Point) GoUp(steps uint) Point {
	return Point{X: p.X, Y: p.Y - int(steps)}
}

func (p Point) GoDown(steps uint) Point {
	return Point{X: p.X, Y: p.Y + int(steps)}
}

func (p Point) GoLeft(steps uint) Point {
	return Point{X: p.X - int(steps), Y: p.Y}
}

func (p Point) GoRight(steps uint) Point {
	return Point{X: p.X + int(steps), Y: p.Y}
}
