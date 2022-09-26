package main

import (
	"github.com/Erickype/GoRoutines/routines"
	"sync"
)

func main() {

	wg := &sync.WaitGroup{}

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go routines.ShowGoRoutine(i, wg)
	}

	wg.Wait()
}
