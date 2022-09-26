package main

import (
	"github.com/Erickype/GoRoutines/routines"
	"time"
)

func main() {
	for i := 0; i < 10; i++ {
		go routines.ShowGoRoutine(i)
	}

	time.Sleep(time.Minute)
}
