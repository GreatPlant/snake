package point

type Point struct {
	X int
	Y int
}

func New(x, y int) Point {
	return Point{
		X: x,
		Y: y,
	}
}
