package main

import (
	"fmt"
	"github.com/Erickype/GoRoutines/routines"
	"sync"
	"time"
)

func main() {

	start := time.Now()
	wg := &sync.WaitGroup{}

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go routines.ShowGoRoutine(i, wg)
	}

	wg.Wait()
	duration := time.Since(start)
	fmt.Printf("Duration: %v\n", duration)
}
