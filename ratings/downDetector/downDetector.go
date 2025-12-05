package downdetector

import (
	"context"
	"fmt"
	"net/http"
	"strings"
	"sync"
)

// Result represents the status check result for a URL
type Result struct {
	URL    string
	Status string
}

// CheckStatus checks the status of a URL using context for cancellation/timeout
// It sends the result through a channel and signals completion via WaitGroup
func CheckStatus(ctx context.Context, url string, resultChan chan<- Result, wg *sync.WaitGroup) {
	defer wg.Done() // Signal that this goroutine is done

	// Create a request with context
	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		resultChan <- Result{URL: url, Status: "down"}
		return
	}

	// Make the HTTP request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		resultChan <- Result{URL: url, Status: "down"}
		return
	}
	defer resp.Body.Close()

	// Check status code
	status := "down"
	if resp.StatusCode == 200 {
		status = "up"
	}

	resultChan <- Result{URL: url, Status: status}
}

// Printer prints the status result
func Printer(result Result) {
	if strings.ToLower(result.Status) == "up" {
		fmt.Println("URL : ", result.URL, " Status Up")
	} else {
		fmt.Println("URL : ", result.URL, " Status Down")
	}
}
