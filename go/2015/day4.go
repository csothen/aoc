package aoc

import (
	"crypto/md5"
	"fmt"
	"strings"
)

func Day4() {
	base, hash := "ckczppom", ""
	// part 1
	hash("00000")

	// part 2
	hash("000000")
}

func hash(prefix string) {
	base, hash := "ckczppom", ""
	i := -1
	for !strings.HasPrefix(hash, prefix) {
		i++
		input := fmt.Sprintf("%s%d", base, i)
		hash = fmt.Sprintf("%x", md5.Sum([]byte(input)))
	}

	fmt.Println(i, hash)
}
