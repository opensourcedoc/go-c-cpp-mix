package main

import (
	"fmt"
	"log"
	point "go-c-cpp-mix"
	"runtime"
)

func main() {
	for i := 0; i < 10; i++ {
		p := point.Init(3, 4)
		fmt.Println(p.ToString())
		point.Free(p)
		runtime.GC()

		var mem runtime.MemStats
		runtime.ReadMemStats(&mem)
		log.Println(mem.Alloc)
		log.Println(mem.TotalAlloc)
		log.Println(mem.HeapAlloc)
		log.Println(mem.HeapSys)
	}
}
