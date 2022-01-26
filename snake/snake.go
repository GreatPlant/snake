package snake

import "snake/internal/point"

type Direction int8

const (
	UP Direction = iota
	LEFT
	DOWN
	RIGHT
)

type Snake struct {
	alive            bool
	head             point.Point
	body             []point.Point
	currentDirection Direction
}

func (s *Snake) IsAlive() bool { return s.alive }

func (s *Snake) Kill() { s.alive = false }

func (s *Snake) Length() int { return len(s.body) + 1 }

func (s *Snake) MakeStep() {
	s.moveBody()
	s.moveHead()
}

func (s *Snake) ChangeDirection(new Direction) {
	s.currentDirection = new
}

func (s *Snake) SelfOccuped() bool {
	for _, segment := range s.body {
		if segment == s.head {
			return true
		}
	}
	return false
}

func (s *Snake) Increase(p point.Point) {
	s.body = prepend(s.body, s.head)
	s.head = p
}

func (s *Snake) moveHead() {
	switch s.currentDirection {
	case UP:
		s.head.Y += 1
	case LEFT:
		s.head.X -= 1
	case DOWN:
		s.head.Y -= 1
	case RIGHT:
		s.head.X += 1
	}
}

func (s *Snake) moveBody() {
	for i := len(s.body) - 1; i > 0; i-- {
		s.body[i] = s.body[i-1]
	}
	s.body[0] = s.head
}

func prepend(s []point.Point, p point.Point) []point.Point {
	s = append(s, point.Point{})
	copy(s[1:], s)
	s[0] = p
	return s
}
