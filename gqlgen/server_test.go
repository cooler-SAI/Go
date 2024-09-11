package main

import (
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/99designs/gqlgen/graphql/handler"
	"gqlgen/graph"
)

func TestGraphQLQueryHandler(t *testing.T) {
	// Define a simple GraphQL query
	query := `{"query":"{ __typename }"}`

	req := httptest.NewRequest("POST", "/query", strings.NewReader(query))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()

	// Initialize the GraphQL handler
	handlerServerDefault := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: &graph.Resolver{}}))
	handlerServerDefault.ServeHTTP(w, req)

	resp := w.Result()
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {

		}
	}(resp.Body)

	// Read the response body
	body, _ := io.ReadAll(resp.Body)

	if resp.StatusCode != http.StatusOK {
		t.Errorf("Expected status code %d, but got %d", http.StatusOK, resp.StatusCode)
		t.Errorf("Response body: %s", string(body))
	}

	// Optionally, assert the response content
	expectedResponseSubstring := `"__typename"`
	if !strings.Contains(string(body), expectedResponseSubstring) {
		t.Errorf("Expected response to contain %s, but got %s", expectedResponseSubstring, string(body))
	}
}
