package main

import (
	"fmt"
	point "go-c-cpp-mix"
	"log"
	"runtime"
)

func main() {
	for i := 0; i < 10; i++ {
		p := point.NewPoint(3, 4)
		fmt.Println(p.String())
		point.Delete(p)
		runtime.GC()

		var mem runtime.MemStats
		runtime.ReadMemStats(&mem)
		log.Println(mem.Alloc)
		log.Println(mem.TotalAlloc)
		log.Println(mem.HeapAlloc)
		log.Println(mem.HeapSys)
	}
}
