package main

import (
	"fmt"
)


// StartDispatcher starts the function that distributes the work
func StartDispatcher(nworkers int) {
	workerQueue = make(chan chan request, nworkers)

	for i := 0; i < nworkers; i++ {
		worker := NewWorker(i+1, workerQueue)
		worker.Start()
	}

	go func() {
		for {
			select {
			case work := <-workQueue:
				fmt.Println("Received work requeust")
				go func() {
					worker := <-workerQueue

					fmt.Println("Dispatching work request")
					worker <- work
				}()
			}
		}
	}()
}
