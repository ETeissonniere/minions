package minions

import (
	"sync"
)

type Command struct {
	Instructions func()
}

type Pool struct {
	queue chan Command
	wg    *sync.WaitGroup
}

func birthMinion(q chan Command, wg *sync.WaitGroup) {
	for j := range q {
		j.Instructions()
	}
	wg.Done()
}

// Create a pool with the given number of minions
func Hire(nbMinions int) *Pool {
	var wg sync.WaitGroup
	q := make(chan Command, nbMinions)

	for i := 0; i < nbMinions; i++ {
		wg.Add(1)
		go birthMinion(q, &wg)
	}

	return &Pool{
		queue: q,
		wg:    &wg,
	}
}

// Push a command with some instructions to the pool
func (p *Pool) Execute(c Command) {
	p.queue <- c
}

// Decommission the minions and wait for them to finish their work
func (p *Pool) Terminate() {
	close(p.queue)
	p.wg.Wait()
}
