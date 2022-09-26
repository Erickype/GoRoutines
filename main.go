package main

import "github.com/Erickype/GoRoutines/routines"

func main() {
	for i := 0; i < 10; i++ {
		routines.ShowGoRoutine(i)
	}
}
