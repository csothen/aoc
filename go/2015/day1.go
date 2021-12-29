package aoc

import (
	"fmt"

	"github.com/csothen/aoc/utils"
)

const (
	D1Up   byte = '('
	D1Down byte = ')'
)

func Day1() {
	d := utils.Read("2015d1.txt")
	floor := 0
	pos := -1
	for i, dir := range d {
		if dir == D1Up {
			floor += 1
			continue
		}
		if dir == D1Down {
			if floor == 0 && pos == -1 {
				pos = i + 1
			}
			floor -= 1
			continue
		}
	}

	fmt.Println(floor, pos)
}
