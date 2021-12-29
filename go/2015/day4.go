package aoc

import (
	"crypto/md5"
	"fmt"
	"strings"
)

func Day4() {
	base, hash := "ckczppom", ""
	i := -1
	for !strings.HasPrefix(hash, "000000") {
		i++
		input := fmt.Sprintf("%s%d", base, i)
		hash = fmt.Sprintf("%x", md5.Sum([]byte(input)))
	}

	fmt.Println(i, hash)
}
