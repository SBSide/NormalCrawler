package main

import (
	"fmt"
	"main/core"
)

func main() {
	var jobs []core.ExtractedJob
	totalpages := core.GetPages()
	fmt.Println(totalpages)
	for i := 0; i < totalpages; i++ {
		extractedJobs := core.GetPage(i)
		jobs = append(jobs, extractedJobs...)
	}
	fmt.Println(jobs)
}
