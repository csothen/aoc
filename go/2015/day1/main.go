package day1

import (
	"fmt"

	"github.com/csothen/aoc/utils"
)

const (
	Up   byte = '('
	Down byte = ')'
)

func Exec() {
	d := utils.Read("2015d1.txt")
	floor := 0
	pos := -1
	for i, dir := range d {
		if dir == Up {
			floor += 1
			continue
		}
		if dir == Down {
			if floor == 0 && pos == -1 {
				pos = i + 1
			}
			floor -= 1
			continue
		}
	}

	fmt.Println(floor, pos)
}
