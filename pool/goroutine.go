package pool

import (
	"sync"
)

type goroutine struct {
	ch chan struct{}
	wg *sync.WaitGroup
}

func (g *goroutine) Add(delta int) {
	g.wg.Add(delta)
	for i := 0; i < delta; i++ {
		g.ch <- struct{}{}
	}
}

func (g *goroutine) Done() {
	<-g.ch
	g.wg.Done()
}

func (g *goroutine) Wait() {
	g.wg.Wait()
}

// NewGoroutine ...
func NewGoroutine(count int) *goroutine {
	return &goroutine{
		ch: make(chan struct{}, count),
		wg: new(sync.WaitGroup),
	}
}
