package day4

import (
	"crypto/md5"
	"fmt"
	"strings"
)

func Exec() {
	base := "ckczppom"
	// part 1
	hashFn(base, "00000")

	// part 2
	hashFn(base, "000000")
}

func hashFn(base, prefix string) {
	hash := ""
	i := -1
	for !strings.HasPrefix(hash, prefix) {
		i++
		input := fmt.Sprintf("%s%d", base, i)
		hash = fmt.Sprintf("%x", md5.Sum([]byte(input)))
	}

	fmt.Println(i, hash)
}
