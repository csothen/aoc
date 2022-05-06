package day2

import (
	"fmt"
	"math"
	"strconv"
	"strings"

	"github.com/csothen/aoc/utils"
)

type Surface struct {
	dim1 int
	dim2 int
	area int
}

func Exec() {
	d := utils.Read("2015d2.txt")
	presents := strings.Split(string(d), "\n")

	totalPaper, totalRibbon := 0, 0
	for _, present := range presents {
		measures := strings.Split(present, "x")
		if len(measures) != 3 {
			continue
		}

		l, _ := strconv.Atoi(measures[0])
		w, _ := strconv.Atoi(measures[1])
		h, _ := strconv.Atoi(measures[2])

		surfaces := []Surface{{l, w, l * w}, {w, h, w * h}, {h, l, h * l}}
		smallestSurface := min(surfaces)

		paper := 2*surfaces[0].area + 2*surfaces[1].area + 2*surfaces[2].area + smallestSurface.area
		totalPaper += paper

		ribbon := 2*smallestSurface.dim1 + 2*smallestSurface.dim2 + (l * w * h)
		totalRibbon += ribbon
	}

	fmt.Println(totalPaper, totalRibbon)
}

func min(surfaces []Surface) Surface {
	min := Surface{area: math.MaxInt32}

	for _, s := range surfaces {
		area := s.dim1 * s.dim2
		if area < min.area {
			min = s
		}
	}
	return min
}
