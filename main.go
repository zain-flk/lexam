package main

import (
	"bytes"
	"fmt"
	"langexam/config"
	"langexam/queue"
	"langexam/r2"
	"log"
	"os/exec"
	"sync"
	"time"

	"github.com/aws/aws-sdk-go-v2/service/s3"
)

const (
	numWorkers = 4 // Number of concurrent workers - TODO: Make it configurable
	duration   = time.Second
)

func executeCommand(command string) {

	cmd := exec.Command("sh", "-c", command)

	var outBuf, errBuf bytes.Buffer
	cmd.Stdout = &outBuf
	cmd.Stderr = &errBuf

	err := cmd.Run()
	if outBuf.Len() > 0 {
		fmt.Println("Output:", outBuf.String())
	}
	if errBuf.Len() > 0 {
		fmt.Println("Error Output:", errBuf.String())
	}

	if err != nil {
		fmt.Println("Execution Error:", err)
	}
}

func worker(id int, jobs <-chan queue.QueueData, wg *sync.WaitGroup, client *s3.Client) {

	for job := range jobs {
		startTime := time.Now() // Start job timer

		err := r2.DownloadFile(client, config.BUCKET_NAME, job.FileName, config.DEST_PATH)
		if err != nil {
			log.Fatalf("Error downloading %s file: %v", job.FileName, err)
		}

		wranglerConfig := fmt.Sprintf("%s/wrangler.jsonc", job.ProjectName)
		wranglerCmd1 := fmt.Sprintf("cp -r langexam %s", job.ProjectName)
		executeCommand(wranglerCmd1)

		config.GenerateWranglerConfig(wranglerConfig, job.ProjectName)
		wranglerCmd2 := fmt.Sprintf("cp downloads/%s %s/src/index.ts && cd %s && npm install && npx wrangler build && npx wrangler deploy && rm -rf %s", job.FileName, job.ProjectName, job.ProjectName, job.ProjectName)
		executeCommand(wranglerCmd2)

		elapsed := time.Since(startTime) // Calculate job time
		fmt.Printf("Worker %d completed task: %s in %v\n", id, job.ProjectName, elapsed)
		wg.Done()
	}
}

func main() {
	config.Init()

	startTime := time.Now() // Start total execution timer
	fmt.Println("Read data from Queue")
	downloads, err := queue.ReadQueue("queue.json")
	if err != nil {
		log.Fatalf("Error: %v", err)
	}
	numExecutions := len(downloads)

	fmt.Println("Connect to file storage")
	client, err := r2.CreateS3Client()
	if err != nil {
		log.Fatalf("Error creating S3 client: %v", err)
	}

	var wg sync.WaitGroup
	jobQueue := make(chan queue.QueueData, numExecutions)

	numWorkers := 10 // min(numExecutions/10, 100) // Dynamically determine the number of workers
	for i := 0; i < numWorkers; i++ {
		go worker(i, jobQueue, &wg, client)
	}

	ticker := time.NewTicker(duration / time.Duration(numExecutions))
	defer ticker.Stop()

	for _, file := range downloads {
		<-ticker.C
		wg.Add(1)
		jobQueue <- file
	}

	wg.Wait() // Wait for all jobs to finish
	close(jobQueue)

	fmt.Println("All commands executed successfully!")
	totalElapsed := time.Since(startTime) // Calculate total execution time
	fmt.Printf("All tasks executed successfully in %v\n", totalElapsed)
}
