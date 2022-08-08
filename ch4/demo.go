package main

import (
	"fmt"
	"strings"
)

type Currency int

const (
	USD Currency = iota
	EUR
	GBP
	RMB
)

func main() {
	//q := [...]int{1, 2, 3}
	//for i := 0; i < len(q); i++ {
	//	fmt.Printf("%d ", q[i])
	//}

	s := "hello world "
	s2 := strings.Join(strings.Fields(s), "")
	fmt.Println(s2)
}
