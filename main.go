package main

import (
	"context"
	"fmt"
	downdetector "hello_world/ratings/downDetector"
	"sync"
	"time"
)

func main() {
	urls := []string{
		"https://gdg.community.dev/gdg-cochin/",
		"https://golang.org",
		"https://httpstat.us/500",
		"https://www.google.com/",
		"https://www.facebook.com/",
		"https://www.twitter.com/",
		"https://www.instagram.com/",
		"https://site-not-present.io",
		"https://www.youtube.com/",
		"https://www.linkedin.com/",
		"https://www.github.com/",
		"https://www.stackoverflow.com/",
		"https://www.reddit.com/",
	}

	// Create a context with 10 second timeout
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// wg is a Synchronization primitive that allows you to wait for a collection of goroutines to finish executing.
	var wg sync.WaitGroup

	// Channel to collect results
	resultChan := make(chan downdetector.Result, len(urls))

	for _, url := range urls {
		wg.Add(1)
		go downdetector.CheckStatus(ctx, url, resultChan, &wg)
	}

	// Channel to signal when all goroutines are done
	done := make(chan struct{})
	go func() {
		wg.Wait()
		close(done)
	}()

	// Wait for either all checks to complete or timeout
	select {
	case <-done:
		fmt.Println("\nAll checks completed!")
		// Process all results
		close(resultChan)
		for result := range resultChan {
			downdetector.Printer(result)
		}
	case <-ctx.Done():
		fmt.Println("\n⏱️  Timeout reached! Some checks may still be running...")
		// Process results received so far (don't close channel as goroutines may still be sending)
		for {
			select {
			case result := <-resultChan:
				downdetector.Printer(result)
			default:
				// No more results available, exit
				return
			}
		}
	}
}
