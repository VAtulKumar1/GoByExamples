package chapters

import (
	"fmt"
	"math/rand"
	"sync/atomic"
	"time"
)

type readOp struct {
	key  int
	rsep chan int
}
type writeOp struct {
	key  int
	val  int
	rsep chan bool
}

func StateFulGoRoutines() {

	var readOps uint64
	var writeOps uint64

	reads := make(chan readOp)
	writes := make(chan writeOp)

	go func() {
		var state = make(map[int]int)
		for {
			select {
			case read := <-reads:
				read.rsep <- state[read.key]
			case write := <-writes:
				state[write.key] = write.val
				write.rsep <- true
			}
		}

	}()

	for r := 0; r < 100; r++ {
		go func() {
			read := readOp{
				key:  rand.Intn(5),
				rsep: make(chan int),
			}

			reads <- read
			<-read.rsep
			atomic.AddUint64(&readOps, 1)
			time.Sleep(time.Millisecond)
		}()
	}

	for w := 0; w < 10; w++ {
		go func() {
			write := writeOp{
				key:  rand.Intn(5),
				val:  rand.Intn(100),
				rsep: make(chan bool),
			}

			writes <- write
			<-write.rsep
			atomic.AddUint64(&writeOps, 1)
			time.Sleep(time.Millisecond)
		}()

	}

	time.Sleep(time.Millisecond)

	readOpsFinal := atomic.LoadUint64(&readOps)
	fmt.Println("readOps:", readOpsFinal)
	writeOpsFinal := atomic.LoadUint64(&writeOps)
	fmt.Println("writeOps:", writeOpsFinal)

}
