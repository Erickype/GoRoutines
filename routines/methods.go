package routines

import (
	"fmt"
	"math/rand"
	"time"
)

func ShowGoRoutine(id int) {
	delay := rand.Intn(500)
	fmt.Printf("Goroutine #%v with %vms delay.\n", id, delay)

	time.Sleep(time.Millisecond * time.Duration(delay))
}
