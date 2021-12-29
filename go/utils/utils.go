package utils

import (
	"fmt"
	"io/ioutil"
	"path"
)

func Read(file string) []byte {
	content, err := ioutil.ReadFile(path.Join("..", "inputs", file))
	if err != nil {
		fmt.Println(err)
		return []byte{}
	}
	return content
}
