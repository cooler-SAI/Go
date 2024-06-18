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
	b.WriteTo(os.Stdout)
}
