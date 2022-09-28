package main

import (
	"fmt"
	"github.com/Erickype/GoRoutines/data"
	"sync"
	"time"
)

func main() {

	start := time.Now()
	wg := &sync.WaitGroup{}
	m := &sync.Mutex{}

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go data.ReadBook(i, wg, m)
	}

	wg.Wait()
	duration := time.Since(start)
	fmt.Printf("Duration: %v\n", duration)
}
