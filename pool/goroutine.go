package pool

import (
	"sync"
)

type goroutine struct {
	c  chan struct{}
	wg *sync.WaitGroup
}

func (g *goroutine) Add(delta int) {
	g.wg.Add(delta)
	for i := 0; i < delta; i++ {
		g.c <- struct{}{}
	}
}

func (g *goroutine) Done() {
	<-g.c
	g.wg.Done()
}

func (g *goroutine) Wait() {
	g.wg.Wait()
}

// NewGoroutine ...
func NewGoroutine(count int) *goroutine {
	return &goroutine{
		c:  make(chan struct{}, count),
		wg: new(sync.WaitGroup),
	}
}
