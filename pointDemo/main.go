package main

import (
	"go-c-cpp-mix/point"
	"log"
)

func main() {
	p := point.NewPoint(0.0, 0.0)
	q := point.NewPoint(3.0, 4.0)

	dist := point.Distance(p, q)
	if 5 != dist {
		log.Fatal("Wrong distance")
	}

	p.Delete()
	q.Delete()
}
