package main

import (
	"flag"
	"fmt"
	"net/http"
	"time"
)

type request struct {
	Name  string
	Delay time.Duration
}


var (
	workQueue = make(chan request, 100)
	workerQueue chan chan request
	numWorkers = flag.Int("n", 7, "The number of workers to start")
	httpAddr = flag.String("http", "localhost:8000", "Address to listen for HTTP requests on")
)

func main() {
	flag.Parse()

	StartDispatcher(*numWorkers)

	http.HandleFunc("/work", Collector)

	fmt.Println("HTTP server listening on", *httpAddr)
	if err := http.ListenAndServe(*httpAddr, nil); err != nil {
		fmt.Println(err.Error())
	}
}
