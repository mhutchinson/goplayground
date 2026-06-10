package main

import (
	"context"
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"strings"
)

func main() {
	ctx := context.Background()
	c := &http.Client{}
	fmt.Println("hello")
	url := "https://api.transparency.dev/notathing/distributor/snakesonaplane"
	req, err := http.NewRequestWithContext(ctx, "PUT", url, strings.NewReader("there is no spoon"))
	if err != nil {
		slog.Error("Failed to create HTTP request", "error", err)
		os.Exit(1)
	}
	resp, err := c.Do(req)
	if err != nil {
		slog.Error("HTTP request failed", "error", err)
		os.Exit(1)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		slog.Error("Bad status response", "status", resp.StatusCode)
		os.Exit(1)
	}
	if resp.Request.Method != "PUT" {
		slog.Error("Request redirect/method conversion", "expected", "PUT", "got", resp.Request.Method, "url", url, "targetURL", resp.Request.URL.String())
		os.Exit(1)
	}
	slog.Info("Request completed successfully", "method", resp.Request.Method)
}
