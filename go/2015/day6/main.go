package day6

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/csothen/aoc/utils"
)

func Exec() {
	size := 1000
	gridP1 := make([][]bool, size)
	gridP2 := make([][]int, size)
	for i := 0; i < size; i++ {
		gridP1[i] = make([]bool, size)
		gridP2[i] = make([]int, size)
		for j := 0; j < size; j++ {
			gridP1[i][j] = false
			gridP2[i][j] = 0
		}
	}

	d := utils.Read("2015d6.txt")
	lines := strings.Split(string(d), "\n")
	for _, line := range lines {
		coords1, coords2, command := getLightsInfo(line)
		gridP1 = processLightsInfoP1(gridP1, coords1, coords2, command)
		gridP2 = processLightsInfoP2(gridP2, coords1, coords2, command)
	}

	count := 0
	brightness := 0
	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			if gridP1[i][j] {
				count++
			}
			brightness += gridP2[i][j]
		}
	}
	fmt.Println("Number of lit lights ->", count)
	fmt.Println("Total brightness ->", brightness)
}

func processLightsInfoP1(grid [][]bool, c1, c2 []int, command string) [][]bool {
	for i := c1[0]; i <= c2[0]; i++ {
		for j := c1[1]; j <= c2[1]; j++ {
			switch command {
			case "toggle":
				grid[i][j] = !grid[i][j]
			case "turn off":
				grid[i][j] = false
			case "turn on":
				grid[i][j] = true
			}
		}
	}
	return grid
}

func processLightsInfoP2(grid [][]int, c1, c2 []int, command string) [][]int {
	for i := c1[0]; i <= c2[0]; i++ {
		for j := c1[1]; j <= c2[1]; j++ {
			switch command {
			case "toggle":
				grid[i][j] += 2
			case "turn off":
				if grid[i][j] == 0 {
					continue
				}
				grid[i][j] -= 1
			case "turn on":
				grid[i][j] += 1
			}
		}
	}
	return grid
}

func getLightsInfo(line string) ([]int, []int, string) {
	split := strings.Split(line, " ")
	command := strings.Join(split[0:len(split)-3], " ")
	c1s := split[len(split)-3]
	c2s := split[len(split)-1]

	coords1 := make([]int, 2)
	coords2 := make([]int, 2)

	c1ss := strings.Split(c1s, ",")
	c2ss := strings.Split(c2s, ",")

	coords1[0], _ = strconv.Atoi(c1ss[0])
	coords1[1], _ = strconv.Atoi(c1ss[1])
	coords2[0], _ = strconv.Atoi(c2ss[0])
	coords2[1], _ = strconv.Atoi(c2ss[1])

	return coords1, coords2, command
}
