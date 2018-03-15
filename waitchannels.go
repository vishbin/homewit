package main

import (
	"fmt"
	"time"
	"sync"
)

func worker(finished chan bool) {
	fmt.Println("Worker: Started")
	time.Sleep(time.Second)
	fmt.Println("Worker: Finished")
	finished <- true
}

func main1() {
	finished := make(chan bool)

	fmt.Println("Main: Starting worker")
	go worker(finished)

	fmt.Println("Main: Waiting for worker to finish")
	<- finished
	fmt.Println("Main: Completed")
}

func worker2(wg *sync.WaitGroup, id int) {
	defer wg.Done()

	fmt.Printf("Worker %v: Started\n", id)
	time.Sleep(time.Second)
	fmt.Printf("Worker %v: Finished\n", id)
}

func main2() {
	var wg sync.WaitGroup

	for i := 0; i < 5; i++ {
		fmt.Println("Main: Starting worker", i)
		wg.Add(1)
		go worker(&wg, i)
	}

	fmt.Println("Main: Waiting for workers to finish")
	wg.Wait()
	fmt.Println("Main: Completed")
}