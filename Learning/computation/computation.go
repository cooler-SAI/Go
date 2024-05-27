package computation

import "fmt"

func ResultNumberNew(n int) string {
	var result int = 1
	result *= n

	return fmt.Sprintf("Result is %d", result)
}
