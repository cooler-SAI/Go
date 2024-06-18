package main

import (
	"bytes"
	"fmt"
	"os"
)

func main() {

	var b bytes.Buffer
	b.WriteString("Hello World")
	fmt.Println(b.String())
	_, _ = b.WriteTo(os.Stdout)
}
