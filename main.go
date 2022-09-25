package main

import (
	"errors"
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func doJob(job *Job, progress *Progress, wg *sync.WaitGroup) {
	defer wg.Done()
	n := rand.Intn(10000)
	time.Sleep(time.Duration(n)*time.Millisecond)
	if rand.Intn(100) > 20 {
		job.Complete(fmt.Sprintf("took %d milliseconds", n))
	} else {
		job.Fail(errors.New("something went wrong"))
	}
	progress.Update()
}

func main() {
	var wg sync.WaitGroup

	//seed := time.Now().UnixNano()
	seed := int64(1664121548889467000)
	fmt.Println("Seed:", seed)
	rand.Seed(seed)

	progress := NewProgress()

	for i := 0; i < 20; i++ {
		j := NewJob(fmt.Sprintf("Building image #%d", i+1))
		progress.AddJob(j)
		wg.Add(1)
		go doJob(j, progress, &wg)
	}

	progress.Start()

	wg.Wait()

	progress.Stop()
}
