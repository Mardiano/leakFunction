package main

import (
	"fmt"
	"log"
	"net/http"
	_ "net/http/pprof"
	leak "profiling-golang/leakFunction"
	"sync"
)

func main() {
	// we need a webserver to get the pprof webserver
	go func() {
		log.Println(http.ListenAndServe("localhost:6060", nil))
	}()
	fmt.Println("hello world")
	var wg sync.WaitGroup
	wg.Add(1)
	go leak.LeakyFunction(wg)
	wg.Wait()
}
