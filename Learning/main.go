package main

import (
	"bytes"
	"fmt"
)

func main() {
	var buf bytes.Buffer
	buf.WriteString("Hello ")
	buf.WriteByte('W')
	buf.WriteString("orld!")

	data := buf.Bytes()
	fmt.Println(string(data))
}
