package main

import (
	"bytes"
	"fmt"
	"os"
)

func main() {

	var byteString bytes.Buffer
	byteString.Write([]byte("Hello"))
	_, _ = fmt.Fprintf(&byteString, "World")
	byteString.WriteTo(os.Stdout)
}
