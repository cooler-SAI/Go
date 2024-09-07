package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"time"

	"github.com/gojek/heimdall/v7/httpclient"
)

func loggingMiddleware(next http.RoundTripper) http.RoundTripper {
	return roundTripperFunc(func(req *http.Request) (*http.Response, error) {
		log.Println("Send Request:", req.URL)

		res, err := next.RoundTrip(req)

		if err != nil {
			log.Println("Error Request:", err)
		} else {
			log.Println("Got Answer:", res.Status)
		}

		return res, err
	})
}

type roundTripperFunc func(*http.Request) (*http.Response, error)

func (f roundTripperFunc) RoundTrip(req *http.Request) (*http.Response, error) {
	return f(req)
}

func main() {
	timeout := 10 * time.Second

	httpClient := &http.Client{
		Transport: loggingMiddleware(http.DefaultTransport),
		Timeout:   timeout,
	}

	client := httpclient.NewClient(
		httpclient.WithHTTPClient(httpClient),
		httpclient.WithRetryCount(3),
	)

	res, err := client.Get("https://api.github.com/repos/gojek/heimdall", nil)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		defer func(Body io.ReadCloser) {
			err := Body.Close()
			if err != nil {
				fmt.Println("Error closing body:", err)
			}
		}(res.Body)
		fmt.Println("Status:", res.StatusCode)
	}
}
