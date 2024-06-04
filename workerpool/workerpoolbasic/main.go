package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type Job struct {
	id       int
	randomno int
}
type Result struct {
	job        Job
	sumofdigit int
}

var jobs = make(chan Job, 10)
var results = make(chan Result, 10)

func sumdigit(number int) int {
	sum := 0
	n := number
	for n > 0 {
		sum += n % 10
		n /= 10

	}
	time.Sleep(1 * time.Millisecond)
	return sum
}
func allocate(noofjob int) {
	for j := 0; j <= noofjob; j++ {
		jobs <- Job{j, rand.Intn(999)}
	}
	//here we close channel job becauase after noofjobs we donot have any job
	close(jobs)
}
func finalresult(done chan bool) {
	for r := range results {
		fmt.Println(r.job.id, r.job.randomno, r.sumofdigit)
	}
	done <- true
}
func worker(wg *sync.WaitGroup) {
	// combine result with job
	for j := range jobs {
		results <- Result{j, sumdigit(j.randomno)}
	}
	wg.Done()

	// we can also close results channel here but doing this means we are clsong channel even there are no of workers available.
}
func createworkerpool(noofworker int) {
	var wg sync.WaitGroup
	for i := 0; i < noofworker; i++ {
		wg.Add(1)
		go worker(&wg)
	}
	wg.Wait()
	//here we close results chanenl because before that all jobs are executes we donot have any jobs remaining so we are informing results function ,
	// functiomn from where we are printing result that results channel have  non value left
	close(results)
}
func main() {
	start := time.Now()
	noofjob := 100
	go allocate(noofjob)
	done := make(chan bool)
	go finalresult(done)
	noofworker := 30
	go createworkerpool(noofworker)

	<-done
	totaltime := time.Since(start)
	fmt.Println("total time taken --", totaltime)

}
