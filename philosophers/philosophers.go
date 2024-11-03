package main

import (
	"fmt"
	"sync"
)

type ChopStick struct{ sync.Mutex }

type Philosopher struct {
	index           int
	leftCS, rightCS *ChopStick
}

func (p Philosopher) eat(host chan bool, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < 3; i++ {
		host <- true
		p.leftCS.Lock()
		p.rightCS.Lock()

		fmt.Println("starting to eat", p.index)
        fmt.Println("finishing eating", p.index)

		p.rightCS.Unlock()
		p.leftCS.Unlock()
		<-host
	}
}

func main() {
	var wg sync.WaitGroup
	host := make(chan bool, 2)

	CSticks := initCSticks()
	philos := initPhilos(CSticks)

	for i := 0; i < 5; i++ {
		wg.Add(1)
		go philos[i].eat(host, &wg)
	}
	wg.Wait()
}

func initPhilos(CSticks []*ChopStick) []*Philosopher {
	philos := make([]*Philosopher, 5)
	for i := 0; i < 5; i++ {
		philos[i] = &Philosopher{i + 1, CSticks[i], CSticks[(i+1)%5]}
	}
	return philos
}

func initCSticks() []*ChopStick {
	CSticks := make([]*ChopStick, 5)
	for i := 0; i < 5; i++ {
		CSticks[i] = new(ChopStick)
	}
	return CSticks
}
