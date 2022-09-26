package routines

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func ShowGoRoutine(id int, wg *sync.WaitGroup) {
	delay := rand.Intn(500)
	fmt.Printf("Goroutine #%v with %vms delay.\n", id, delay)

	time.Sleep(time.Millisecond * time.Duration(delay))

	wg.Done()
}
