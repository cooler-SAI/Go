package main

import (
	"bytes"
	"fmt"
	"github.com/rs/zerolog"
	"testing"
)

func setupTestLogger() *bytes.Buffer {
	var buf bytes.Buffer
	Logger = zerolog.New(&buf).With().Timestamp().Logger()
	return &buf
}

func TestIntMinBasic(t *testing.T) {
	buf := setupTestLogger()

	ans := IntMin(2, -2)
	if ans != -2 {
		t.Errorf("IntMin(2, -2) = %d; want -2", ans)
	}

	t.Logf("Logs: %s", buf.String())
}

func TestIntMinTableDriven(t *testing.T) {
	buf := setupTestLogger()

	var tests = []struct {
		a, b int
		want int
	}{
		{0, 1, 0},
		{1, 0, 0},
		{2, -2, -2},
		{0, -1, -1},
		{-1, 0, -1},
	}

	for _, tt := range tests {
		testName := fmt.Sprintf("%d,%d", tt.a, tt.b)
		t.Run(testName, func(t *testing.T) {
			ans := IntMin(tt.a, tt.b)
			if ans != tt.want {
				t.Errorf("got %d, want %d", ans, tt.want)
			}
		})
	}

	t.Logf("Logs: %s", buf.String())
}

func BenchmarkIntMin(b *testing.B) {
	buf := setupTestLogger()

	for i := 0; i < b.N; i++ {
		IntMin(1, 2)
	}

	b.Logf("Logs: %s", buf.String())
}

func TestHello(t *testing.T) {
	buf := setupTestLogger()

	got := Hello()
	want := "Hello Tester!"
	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}

	t.Logf("Logs: %s", buf.String())
}

func TestResult(t *testing.T) {
	buf := setupTestLogger()

	tests := []struct {
		a, b     int
		expected int
	}{
		{10, 5, 5},
		{5, 10, -5},
		{10, 10, 0},
	}

	for _, tt := range tests {
		actual := Result(tt.a, tt.b)
		if actual != tt.expected {
			t.Errorf("Result(%d, %d) = %d; expected %d", tt.a, tt.b, actual, tt.expected)
		}
	}

	t.Logf("Logs: %s", buf.String())
}
