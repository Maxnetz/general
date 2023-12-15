package main

import (
	"fmt"
)

// jobs
type number struct {
	a int
	b int
}

// results
type sum struct {
	r int
}

func main() {
	// Basic -> sync.WaitGroup
	// a := []int{1, 2, 3, 4, 5}

	// Basic for loop
	// for i := range a {
	// 	fmt.Printf("i: %v", a[i])
	// }

	// WaitGroup
	// var wg sync.WaitGroup
	// for i := range a {
	// 	wg.Add(1)
	// 	go func(i int) {
	// 		defer wg.Done()
	// 		fmt.Printf("%v ", a[i])
	// 	}(i)
	// }
	// wg.Wait() // result: 3 1 5 4 2

	// Channels
	// numberCh := make(chan int)
	// msgCh := make(chan string)

	// // number
	// go func(numberCh chan<- int) {
	// 	numberCh <- 10
	// }(numberCh)

	// // msg
	// go func(msgCh chan<- string) {
	// 	msgCh <- "Hello World!"
	// }(msgCh)

	// number := <-numberCh
	// msg := <-msgCh

	// fmt.Println()
	// fmt.Println(number)
	// fmt.Println(msg)

	// Worker Pool
	num := []number{
		{a: 1, b: 2},
		{a: 2, b: 3},
		{a: 3, b: 4},
		{a: 4, b: 5},
		{a: 5, b: 6},
		{a: 1, b: 2},
		{a: 2, b: 3},
		{a: 3, b: 4},
		{a: 4, b: 5},
		{a: 5, b: 6},
	}

	jobsCh := make(chan number, len(num))
	resultsCh := make(chan sum, len(num))

	for _, n := range num {
		jobsCh <- n
	}
	close(jobsCh)

	numWorkers := 2
	for w := 0; w < numWorkers; w++ {
		go workers(jobsCh, resultsCh)
	}

	results := make([]sum, 0)
	for a := 0; a < len(num); a++ {
		temp := <-resultsCh
		results = append(results, temp)
	}

	fmt.Println(results)
}

func workers(jobsCh <-chan number, resultsCh chan<- sum) {
	for job := range jobsCh {
		resultsCh <- sum{summation(job.a, job.b)}
	}
}

func summation(a, b int) int {
	return a + b
}
