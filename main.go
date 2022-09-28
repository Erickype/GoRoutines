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

	for i := 1; i <= 5; i++ {
		wg.Add(1)
		go data.ReadBook(i, wg)
	}

	wg.Wait()
	duration := time.Since(start)
	fmt.Printf("Duration: %v\n", duration)
}
