package main

import (
	"bytes"
	"fmt"
)

func main() {

	stringBuffer := new(bytes.Buffer)
	stringBuffer.WriteString("Hello ")
	stringBuffer.WriteString("World")
	fmt.Println("Result of Buffer is: ", stringBuffer.String())
}
