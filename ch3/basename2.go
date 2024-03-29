package main

import (
	"fmt"
	"strings"
)

func basename2(s string) string {
	slash := strings.LastIndex(s, "/")
	s = s[slash+1:]
	if dot := strings.LastIndex(s, "."); dot >= 0 {
		s = s[:dot]
	}
	return s
}

func main() {
	s := "C:/temp/hello.txt"
	fmt.Printf(basename2(s))
}
