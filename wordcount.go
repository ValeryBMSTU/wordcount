package main

import (
	"fmt"
	"strings"
	"os"
)

func main() {
	src := os.Args[1]
	if len(src) == 0 {
		fmt.Println(0)
		return
	}
	srcLen := len(strings.Split(src, " "))
	fmt.Println(srcLen)
}
