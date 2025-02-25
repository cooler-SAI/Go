package misc

import (
	"fmt"
	"sync"
	"time"
)

func sender(id int, message <-chan string, results chan<- string, wg *sync.WaitGroup) {
	defer wg.Done()
	for message := range message {
		fmt.Printf("Sender %d sent message: %s\n", id, message)
		time.Sleep(time.Millisecond * 500)
		results <- fmt.Sprintf("Sender %d sucsessful sent: %s", id, message)
	}
}

func main() {

	messagesList := []string{
		"Hello",
		"All Here",
		"Hello Henry",
		"Hello Ann",
		"Hello Kristy",
		"Hello Mike",
	}

	messagesChan := make(chan string, len(messagesList))
	resultsChan := make(chan string, len(messagesList))

	var wg sync.WaitGroup

	numSenders := 2

	for i := 1; i <= numSenders; i++ {
		wg.Add(1)
		go sender(i, messagesChan, resultsChan, &wg)
	}

	for _, msg := range messagesList {
		messagesChan <- msg
	}
	close(messagesChan)

	wg.Wait()
	close(resultsChan)

	fmt.Println("Result of sending messages:")
	for res := range resultsChan {
		fmt.Println(res)
	}
}
