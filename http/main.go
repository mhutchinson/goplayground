package main

import (
	"context"
	"fmt"
	"net/http"
	"strings"

	"k8s.io/klog/v2"
)

func main() {
	ctx := context.Background()
	c := &http.Client{}
	fmt.Println("hello")
	url := "https://api.transparency.dev/notathing/distributor/snakesonaplane"
	req, _ := http.NewRequest("PUT", url, strings.NewReader("there is no spoon"))
	resp, err := c.Do(req.WithContext(ctx))
	if err != nil {
		klog.Exit(err)
	}
	if resp.StatusCode != 200 {
		klog.Exitf("Bad status response: %d", resp.StatusCode)
	}
	if resp.Request.Method != "PUT" {
		klog.Exitf("PUT request to %q was converted to %s request to %q", url, resp.Request.Method, resp.Request.URL)
	}
	klog.Info(resp.Request.Method)
}
