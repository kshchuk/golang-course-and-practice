package main

/*
    A race condition in Go (or any concurrent programming language) occurs when two or more goroutines
(concurrent threads of execution in Go) access shared data concurrently, and at least one of them
modifies the data, leading to unpredictable or undesirable behavior. These race conditions can result
in bugs that are often difficult to detect and reproduce because the outcome depends on the timing and
order of execution of the goroutines.

    Race conditions can occur in Go when the following conditions are met:

Shared Data:

    There must be shared data that multiple goroutines access or modify. This shared data can be variables,
    data structures, files, or any resource that can be accessed concurrently.

Lack of Synchronization:

    The goroutines do not use proper synchronization mechanisms to control access to the shared data. In Go,
    synchronization can be achieved using mutexes, channels, or other concurrency primitives.
*/

import (
	"fmt"
	"sync"
)

var counter int
var wg sync.WaitGroup

func main() {
	wg.Add(2)
	go increment()
	go decrement()
	wg.Wait()
	fmt.Println("Final Counter Value:", counter)
}

func increment() {
	for i := 0; i < 1000000; i++ {
		counter++
	}
	wg.Done()
}

func decrement() {
	for i := 0; i < 1000000; i++ {
		counter--
	}
	wg.Done()
}

/*
The race condition occurs because both goroutines are accessing and modifying the counter
variable simultaneously without any synchronization mechanism to protect it.
As a result, the final value of counter is unpredictable, and running the program multiple
times may yield different results.
*/
