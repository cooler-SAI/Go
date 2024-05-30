package main

import (
	"fmt"
	"io"
)

type phoneReader string
type phoneWriter struct {
}

func (ph phoneReader) Read(p []byte) (int, error) {
	count := 0
	for i := 0; i < len(ph); i++ {
		if ph[i] >= '0' && ph[i] <= '9' {
			p[count] = ph[i]
			count++
		}
	}
	return count, io.EOF
}

func (ph phoneWriter) Write(bs []byte) (int, error) {
	if len(bs) == 0 {
		return 0, nil
	}
	for i := 0; i < len(bs); i++ {
		if bs[i] >= '0' && bs[i] <= '9' {
			fmt.Print(string(bs[i]))

		}
	}
	fmt.Println()
	return len(bs), nil
}

func main() {

	phone1 := phoneReader("+3-632-78-124-20")
	phone2 := phoneReader("+2-52-345-856-19")

	buffer := make([]byte, len(phone1))
	phone1.Read(buffer)
	fmt.Println(string(buffer))

	buffer = make([]byte, len(phone2))
	phone2.Read(buffer)
	fmt.Println(string(buffer))

	bytes1 := []byte("+123456789")
	bytes2 := []byte("+987654321")

	writer := phoneWriter{}
	writer.Write(bytes1)
	writer.Write(bytes2)

}
