package main

import (
	"bytes"
	"fmt"
	"log"
	"os"
)

func main() {

	log.SetFlags(log.LstdFlags | log.Lmicroseconds)
	log.Println("with micro")

	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Println("with file/line")

	logger := log.New(os.Stdout, "my:", log.LstdFlags)
	logger.Println("from logger")

	var buf bytes.Buffer
	bufLog := log.New(&buf, "buf:", log.LstdFlags)
	bufLog.Println("hello")
	fmt.Print("from bufLog:", buf.String())

}
