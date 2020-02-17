package main

import (
	"log"
	"os"
	"sync"
	"time"
)

func main() {
	log.Println("start")
	start := time.Now()
	solve(os.Args[1])
	log.Printf("done after %v", time.Since(start).Round(time.Millisecond))
}

func solve(name string) {
	model := ModelFromFile(FileFrom(name))

	var wg sync.WaitGroup
	run := func(name string, fn func(Model)) {
		log.Println(name)
		start := time.Now()
		fn(model)
		log.Printf("%v done after %v", name, time.Since(start).Round(time.Millisecond))
		wg.Done()
	}
	wg.Add(2)
	go run("evol", runEvol)
	go run("greedy", solveGreedy)
	wg.Wait()
}
