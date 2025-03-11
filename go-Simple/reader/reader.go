package reader

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func Reader(filename string) ([]string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, fmt.Errorf("can't open file: %w", err)
	}

	defer func(file *os.File) {
		if err := file.Close(); err != nil {
			fmt.Println("Can't close file, sorry...", err)
		}
	}(file)

	scanner := bufio.NewScanner(file)
	var filteredNames []string
	for scanner.Scan() {
		name := scanner.Text()
		if strings.HasPrefix(name, "A") {
			filteredNames = append(filteredNames, name)
		}
	}

	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("reading file: %w", err)
	}

	return filteredNames, nil
}
