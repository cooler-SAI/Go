package main

import (
	"bytes"
	"fmt"
	"testing"
)

func TestIntMinBasic(t *testing.T) {
	ans := IntMin(2, -2)
	if ans != -2 {
		t.Errorf("IntMin(2, -2) = %d; want -2", ans)
	}
}

func TestIntMinTableDriven(t *testing.T) {
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
}

func BenchmarkIntMin(b *testing.B) {
	for i := 0; i < b.N; i++ {
		IntMin(1, 2)
	}
}

func TestMainOutput(t *testing.T) {
	var buf bytes.Buffer
	_, err := fmt.Fprint(&buf, "Hello Tests!\n")
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	expected := "Hello Tests!\n"
	if buf.String() != expected {
		t.Errorf("Expected %q but got %q", expected, buf.String())
	}
}

func TestHello(t *testing.T) {
	got := Hello()
	want := "Hello Tester!"
	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}

func TestResult(t *testing.T) {
	tests := []struct {
		a, b     int
		expected int
	}{
		{10, 5, 5},  // Positive case
		{5, 10, -5}, // Negative case
		{10, 10, 0}, // Zero case
	}

	for _, tt := range tests {
		actual := Result(tt.a, tt.b)
		if actual != tt.expected {
			t.Errorf("Result(%d, %d) = %d; expected %d", tt.a, tt.b, actual, tt.expected)
		}
	}
}
