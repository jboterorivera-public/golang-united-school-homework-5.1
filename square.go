package square

type Point struct {
	x, y int
}

type Square struct {
	start Point
	a     uint
}

func (square Square) End() Point {
	// implement me
	endPoint := Point{x: square.start.x + int(square.a), y: square.start.y + int(square.a)}

	return endPoint
}

func (square Square) Area() uint {
	// implement me
	return uint(square.a * square.a)
}

func (square Square) Perimeter() uint {
	// implement me
	return uint(square.a * 4)
}
