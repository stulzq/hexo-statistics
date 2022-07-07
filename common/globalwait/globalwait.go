package globalwait

import (
	"os"
	"os/signal"
	"sync"
)

var wg = &sync.WaitGroup{}

func Add(delta int) {
	wg.Add(delta)
}

func Done() {
	wg.Done()
}

func Wait() {
	wg.Wait()
}

func NewInterrupt() chan os.Signal {
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt)
	return interrupt
}
