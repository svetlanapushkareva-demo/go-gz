package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Калькулятор")

	operation := scanUserInput()
	fmt.Println("Вы ввели операцию:", operation)

	numbers := getNumbers()
	fmt.Println("Введенные числа:", numbers)

	switch strings.ToUpper(operation) {
	case "SUM":
		result := sum(numbers)
		fmt.Printf("Сумма: %.2f\n", result)
	case "AVG":
		result := avg(numbers)
		fmt.Printf("Среднее: %.2f\n", result)
	case "MED":
		result := med(numbers)
		fmt.Printf("Медиана: %.2f\n", result)
	default:
		fmt.Println("Неизвестная операция. Используйте: SUM, AVG или MED")
	}
}

func scanUserInput() string {
	var input string
	for {
		fmt.Println("Введите операцию (AVG, SUM, MED):")
		fmt.Scan(&input)

		if input != "AVG" && input != "SUM" && input != "MED" {
			fmt.Println("Вы ввели неверное значение")
			continue
		}

		return input
	}
}

func getNumbers() []float64 {
	var input string
	fmt.Scan(&input)

	parts := strings.Split(input, ",")
	var numbers []float64

	for _, part := range parts {
		part = strings.TrimSpace(part)
		if num, err := strconv.ParseFloat(part, 64); err == nil {
			numbers = append(numbers, num)
		}
	}

	return numbers
}

func sum(numbers []float64) float64 {
	sum := 0.0
	for _, num := range numbers {
		sum += num
	}
	return sum
}

func avg(numbers []float64) float64 {
	average := sum(numbers) / float64(len(numbers))
	return average
}

func med(numbers []float64) float64 {
	sorted := make([]float64, len(numbers))
	copy(sorted, numbers)
	sort.Float64s(sorted)

	var median float64
	if len(sorted)%2 == 1 {
		median = sorted[len(sorted)/2]
	} else {
		mid1 := sorted[len(sorted)/2-1]
		mid2 := sorted[len(sorted)/2]
		median = (mid1 + mid2) / 2
	}

	return median
}
