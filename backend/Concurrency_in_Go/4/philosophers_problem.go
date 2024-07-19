package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

// Chopstick is a mutex that represents a chopstick
type Chopstick struct {
	sync.Mutex
}

// Philosopher is a struct with an id, two chopsticks,
// and a channel to communicate with the host
type Philosopher struct {
	id             int
	leftChopstick  *Chopstick
	rightChopstick *Chopstick
	eatCount       int
}

// Host is a struct with a slice of philosophers
type Host struct {
	philosophers []*Philosopher
}

// NewHost creates a new host with a slice of philosophers
func NewHost(philosophers []*Philosopher) *Host {
	return &Host{philosophers: philosophers}
}

// NewPhilosopher creates a new philosopher with an id and two chopsticks
func NewPhilosopher(id int, leftChopstick, rightChopstick *Chopstick) *Philosopher {
	return &Philosopher{
		id:             id,
		leftChopstick:  leftChopstick,
		rightChopstick: rightChopstick,
	}
}

// StartEating starts the eating process for a philosopher
func (p *Philosopher) Eat() {
	p.leftChopstick.Lock()
	p.rightChopstick.Lock()

	fmt.Println("starting to eat", p.id)
	time.Sleep(time.Second)
	p.eatCount++

	p.rightChopstick.Unlock()
	p.leftChopstick.Unlock()
	fmt.Println("finishing eating", p.id)
}

// StartEating starts the eating process for a host
// Each philosopher should eat 3 times
// The philosophers pick up the chopstick in any order
// The host allows no more than 2 philosophers to eat concurrently
func (h *Host) StartEating() {
	wg.Add(1)
	// Create a channel to limit concurrent access to philosophers
	// Set its capacity to 2 to allow two philosophers to eat at a time
	philosopherLimit := make(chan struct{}, 2)

	// Create a WaitGroup to wait for all philosophers to finish eating
	var philosophersFinished sync.WaitGroup

	// Loop through each philosopher and let them eat 3 times
	for _, philosopher := range h.philosophers {
		philosopher := philosopher // Capture the current philosopher for the goroutine

		// Start a goroutine for each philosopher
		go func() {
			for i := 0; i < 3; i++ { // Each philosopher should eat 3 times
				philosopherLimit <- struct{}{} // Acquire access to philosopher
				philosopher.Eat()              // Philosopher starts eating
				<-philosopherLimit             // Release access to philosopher
			}
			philosophersFinished.Done()
		}()
		philosophersFinished.Add(1)
	}

	// Wait for all philosophers to finish eating
	philosophersFinished.Wait()
	wg.Done()
}

func main() {
	// Create 5 chopsticks
	chopsticks := make([]*Chopstick, 5)
	for i := 0; i < 5; i++ {
		chopsticks[i] = &Chopstick{}
	}

	// Create 5 philosophers
	philosophers := make([]*Philosopher, 5)
	for i := 0; i < 5; i++ {
		philosophers[i] = NewPhilosopher(i+1, chopsticks[i], chopsticks[(i+1)%5])
	}

	// Create a host
	host := NewHost(philosophers)

	// Start eating
	go host.StartEating()

	// Wait for the philosophers to start eating
	time.Sleep(time.Millisecond)

	// Wait for the philosophers to finish eating
	wg.Wait()
}
