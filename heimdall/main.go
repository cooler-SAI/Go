package main

import (
	"fmt"
	"io"
	"time"

	"github.com/gojek/heimdall/v7/httpclient"
)

func main() {

	timeout := 10 * time.Second

	client := httpclient.NewClient(httpclient.WithHTTPTimeout(timeout), httpclient.WithRetryCount(3))

	res, err := client.Get("https://api.example.com/data", nil)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		defer func(Body io.ReadCloser) {
			err := Body.Close()
			if err != nil {

			}
		}(res.Body)
		fmt.Println("Status:", res.StatusCode)
	}
}
