package main

import (
  "fmt"
  "time"
)

// Worker definition.
type Worker struct {
  ID          int
  Work        chan request
  WorkerQueue chan chan request
  QuitChan    chan bool
}

// NewWorker creates a new worker.
func NewWorker(id int, workerQueue chan chan request) Worker {
  return Worker{
    ID:          id,
    Work:        make(chan request),
    WorkerQueue: workerQueue,
    QuitChan:    make(chan bool),
  }
}

// Start starts the worker.
func (w *Worker) Start() {
    go func() {
      for {
        // Add self to worker queue
				w.WorkerQueue <- w.Work
				fmt.Println("worker ", w.ID, " added to worker queue")

        select {
        case work := <-w.Work:
          fmt.Printf("worker%d: Received work request, delaying for %f seconds\n", w.ID, work.Delay.Seconds())

          time.Sleep(work.Delay)
          fmt.Printf("worker%d: Hello, %s!\n", w.ID, work.Name)

        case <-w.QuitChan:
          fmt.Printf("worker%d stopping\n", w.ID)
          return
        }
      }
    }()
}

// Stop stops the worker.
func (w *Worker) Stop() {
  go func() {
    w.QuitChan <- true
  }()
}
