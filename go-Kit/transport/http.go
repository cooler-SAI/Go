package transport

import (
	"context"
	"encoding/json"
	"go-Kit/endpoint"
	"go-Kit/service"
	"net/http"

	kitHttp "github.com/go-kit/kit/transport/http"
)

func NewHTTPServer(svc service.HelloService) http.Handler {
	helloHandler := kitHttp.NewServer(
		endpoint.MakeHelloEndpoint(svc),
		decodeHelloRequest,
		encodeResponse,
	)

	r := http.NewServeMux()
	r.Handle("/hello", helloHandler)
	return r
}

func decodeHelloRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var req endpoint.HelloRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	return req, err
}

func encodeResponse(_ context.Context, w http.ResponseWriter, response interface{}) error {
	return json.NewEncoder(w).Encode(response)
}
