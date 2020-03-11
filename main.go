package main

import (
	"fmt"
	"main/core"
)

func main() {
	var jobs []core.ExtractedJob
	c := make(chan []core.ExtractedJob)
	totalpages := core.GetPages()
	for i := 0; i < totalpages; i++ {
		go core.GetPage(i, c)
	}

	for i := 0; i < totalpages; i++ {
		extractedJobs := <-c
		jobs = append(jobs, extractedJobs...)
	}
	core.WriteJobs(jobs)
	fmt.Println("Done, extracted jobs : ", len(jobs))
}
