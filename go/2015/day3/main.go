package day3

import (
	"fmt"

	"github.com/csothen/aoc/utils"
)

const (
	Up    byte = '^'
	Down  byte = 'v'
	Right byte = '>'
	Left  byte = '<'
)

type Point struct {
	x int
	y int
}

func Exec() {
	d := utils.Read("2015d3.txt")
	m := make(map[Point]int)

	x1, x2, y1, y2 := 0, 0, 0, 0
	m[Point{x1, y1}] = 1
	for i, dir := range d {
		var x, y int
		if i%2 == 0 {
			move(dir, &x1, &y1)
			x, y = x1, y1
		} else {
			move(dir, &x2, &y2)
			x, y = x2, y2
		}

		p := Point{x, y}
		m[p] = m[p] + 1
	}

	fmt.Println(len(m))
}

func move(dir byte, x, y *int) {
	if dir == Right {
		*x += 1
	}
	if dir == Down {
		*y += 1
	}
	if dir == Left {
		*x -= 1
	}
	if dir == Up {
		*y -= 1
	}
}
