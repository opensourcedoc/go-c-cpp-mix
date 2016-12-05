package main

import (
	"fmt"
	point "go-c-cpp-mix"
)

func main() {
	p := point.Init(3, 4)
	defer p.Free()
	fmt.Println(p.ToString())
}
