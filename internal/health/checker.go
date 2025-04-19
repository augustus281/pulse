package health

import (
	"fmt"
	"net/http"
	"time"
)

func CheckHealth(url string) error {
	client := &http.Client{
		Timeout: 5 * time.Second,
	}
	resp, err := client.Get(url)
	if err != nil {
		return fmt.Errorf("failed to make GET request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode >= 200 && resp.StatusCode < 300 {
		fmt.Printf("Health check passed: %d\n", resp.StatusCode)
	} else {
		fmt.Printf("Health check failed: %d\n", resp.StatusCode)
	}
	return nil
}
