package misc

import "math"

func CalculatePlinth(totalMeters float64, costPerPlinth float64) (int, float64) {
	plinthLength := 2.7
	plinthCount := int(math.Ceil(totalMeters / plinthLength))
	totalCost := float64(plinthCount) * costPerPlinth
	return plinthCount, totalCost
}

func CheckResult(firstNum int, secondNum int) int {
	result := firstNum + secondNum
	return result
}

func SumAndAverage(numbers []int) (int, float64) {

	if len(numbers) == 0 {
		return 0, 0.0
	}

	sum := 0
	for _, num := range numbers {
		sum += num
	}

	average := float64(sum) / float64(len(numbers))

	return sum, average
}

func ReverseAndCountUnique(s string) (string, int) {

	reversed := ""
	for i := len(s) - 1; i >= 0; i-- {
		reversed += string(s[i])
	}

	uniqueChars := make(map[rune]struct{})
	for _, char := range s {
		uniqueChars[char] = struct{}{}
	}

	return reversed, len(uniqueChars)
}
