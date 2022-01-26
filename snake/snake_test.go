package snake

import (
	"snake/internal/point"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSnakeLength(t *testing.T) {
	snake := Snake{
		head: point.New(1, 2),
		body: []point.Point{point.New(1, 3), point.New(1, 4)},
	}
	assert.Equal(t, 3, snake.Length())
}

func TestSnakeMakeStepDown(t *testing.T) {
	snake := Snake{
		head:             point.New(1, 2),
		body:             []point.Point{point.New(1, 3), point.New(1, 4)},
		currentDirection: DOWN,
	}
	after := Snake{
		head:             point.New(1, 1),
		body:             []point.Point{point.New(1, 2), point.New(1, 3)},
		currentDirection: DOWN,
	}
	snake.MakeStep()
	assert.Equal(t, after, snake)
}

func TestSnakeMakeStepUp(t *testing.T) {
	snake := Snake{
		head:             point.New(1, 4),
		body:             []point.Point{point.New(1, 3), point.New(1, 2)},
		currentDirection: UP,
	}
	after := Snake{
		head:             point.New(1, 5),
		body:             []point.Point{point.New(1, 4), point.New(1, 3)},
		currentDirection: UP,
	}
	snake.MakeStep()
	assert.Equal(t, after, snake)
}

func TestSnakeSelfOccupedTrue(t *testing.T) {
	snake := Snake{
		alive: true,
		head:  point.New(1, 1),
		body:  []point.Point{point.New(0, 0), point.New(0, 2), point.New(1, 2), point.New(1, 1)},
	}
	assert.True(t, snake.SelfOccuped())
}

func TestSnakeSelfOccupedFalse(t *testing.T) {
	snake := Snake{
		alive: true,
		head:  point.New(1, 3),
		body:  []point.Point{point.New(1, 2), point.New(1, 1)},
	}
	assert.False(t, snake.SelfOccuped())
}

func TestSnakeIncrease(t *testing.T) {
	snake := Snake{
		alive: true,
		head:  point.New(1, 4),
		body:  []point.Point{point.New(1, 3), point.New(1, 2)},
	}
	after := Snake{
		alive: true,
		head:  point.New(1, 5),
		body:  []point.Point{point.New(1, 4), point.New(1, 3), point.New(1, 2)},
	}
	snake.Increase(point.New(1, 5))
	assert.Equal(t, after, snake)
}
