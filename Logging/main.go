package main

import (
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

}
