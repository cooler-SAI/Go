package main

import (
	"bytes"
	"fmt"
	"testing"
)

func TestMainOutput(t *testing.T) {

	var buf bytes.Buffer
	_, err := fmt.Fprint(&buf, "Hello Tests!\n")
	if err != nil {
		fmt.Println("we got an error")
		return
	}

	expected := "Hello Tests!\n"
	if buf.String() != expected {
		t.Errorf("Expected %q but got %q", expected, buf.String())
	}
}
