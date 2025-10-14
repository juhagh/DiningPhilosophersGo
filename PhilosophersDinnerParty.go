package main

import (
	"fmt"
	"sync"
)

// Each ChopS represents a lockable chopstick using embedded sync.Mutex
type ChopS struct{ sync.Mutex }

// Philosopher2 models a diner with an ID, two adjacent chopsticks, and access to the host via request channel
type Philosopher struct {
	philosopherID     int
	firstCS, secondCS *ChopS
	reqCh             chan *Request
}

// Request encapsulates a philosopher's permission request to the host, including reply and completion channels
type Request struct {
	id      int
	replyCh chan bool
	doneCh  chan struct{}
}

var wg sync.WaitGroup

func (p *Philosopher) eat() {
	eats := 0
	for eats < 3 { // Guarantee that each philosopher gets to eat exactly 3 times
		replyCh := make(chan bool)
		doneCh := make(chan struct{})
		p.reqCh <- &Request{id: p.philosopherID, replyCh: replyCh, doneCh: doneCh}

		if <-replyCh { // Host granted permission to eat
			p.firstCS.Lock()
			p.secondCS.Lock()

			fmt.Printf("starting to eat %d\n", p.philosopherID)
			//time.Sleep(500 * time.Millisecond)
			fmt.Printf("finishing eating %d\n", p.philosopherID)

			p.firstCS.Unlock()
			p.secondCS.Unlock()

			doneCh <- struct{}{} // Signal to host that this eating round is complete
			eats++
		}
	}
	wg.Done()
}

func host(requestCh <-chan *Request) {
	active := 0
	var mu sync.Mutex

	for req := range requestCh {
		// The host allows no more than 2 philosophers to eat concurrently.
		// Protect active counter with mutex to prevent race conditions
		mu.Lock()
		if active < 2 {
			active++
			req.replyCh <- true

			// Wait for completion signal and decrement active count
			go func(done <-chan struct{}) {
				<-done
				mu.Lock()
				active--
				mu.Unlock()
			}(req.doneCh)
		} else {
			req.replyCh <- false
		}
		mu.Unlock()
	}

}

func main() {
	// Initialise 5 chopsticks
	CSticks := make([]*ChopS, 5)
	for i := 0; i < 5; i++ {
		CSticks[i] = new(ChopS)
	}

	// Shared channel for philosophers to request permission from host
	requestCh := make(chan *Request)

	// Create 5 philosophers (ID 1–5), each with left and right chopsticks
	philosophers := make([]*Philosopher, 5)
	for i := 0; i < 5; i++ {
		philosophers[i] = &Philosopher{i + 1, CSticks[i], CSticks[(i+1)%5], requestCh}
	}

	// Start host routine
	go host(requestCh)

	wg.Add(5)
	// Start the dinner party
	for i := 0; i < 5; i++ {
		go philosophers[i].eat()
	}
	wg.Wait()

}
