package aoc

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/csothen/aoc/utils"
)

var (
	reHasVowels = regexp.MustCompile(`.*([aeiou].*){3}`)
	reIsNotOk   = regexp.MustCompile(`(ab|cd|pq|xy)`)
)

func Day5() {
	d := utils.Read("2015d5.txt")
	words := strings.Split(string(d), "\n")
	nGood1, nGood2 := 0, 0
	for _, w := range words {
		if isOkP1(w) {
			nGood1++
		}
		if isOkP2(w) {
			nGood2++
		}
	}

	fmt.Println("Part 1 -> ", nGood1)
	fmt.Println("Part 2 -> ", nGood2)
}

func isOkP1(word string) bool {
	return reHasVowels.MatchString(word) && !reIsNotOk.MatchString(word) && hasRepeatingLetter(word)
}

func isOkP2(word string) bool {
	return hasRepeatWithLetterBetween(word) && hasRepeatPairWithNoOverlap(word)
}

func hasRepeatPairWithNoOverlap(word string) bool {
	if len(word) < 5 {
		return false
	}

	for i := 0; i < len(word)-1; i++ {
		if strings.Contains(word[i+2:], word[i:i+2]) {
			return true
		}
	}
	return false
}

func hasRepeatWithLetterBetween(word string) bool {
	if len(word) < 3 {
		return false
	}

	for i := 0; i < len(word)-2; i++ {
		if word[i] == word[i+2] {
			return true
		}
	}
	return false
}

func hasRepeatingLetter(word string) bool {
	if len(word) < 2 {
		return false
	}

	for i := 1; i < len(word); i++ {
		if word[i] == word[i-1] {
			return true
		}
	}
	return false
}
