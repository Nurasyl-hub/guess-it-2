package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func mean(numbers []int) float64 {
	sum := 0
	for _, num := range numbers {
		sum += num
	}
	return float64(sum) / float64(len(numbers))
}

func stdDev(numbers []int, meanVal float64) float64 {
	var sumSquares float64
	for _, num := range numbers {
		sumSquares += math.Pow(float64(num)-meanVal, 2)
	}
	return math.Sqrt(sumSquares / float64(len(numbers)))
}

func calculateRange(numbers []int) (int, int) {
	if len(numbers) < 2 {
		return numbers[0] - 10, numbers[0] + 10
	}

	meanVal := mean(numbers)
	stdev := stdDev(numbers, meanVal)

	lowerLimit := meanVal - 1.5*stdev
	upperLimit := meanVal + 1.5*stdev

	return int(math.Round(lowerLimit)), int(math.Round(upperLimit))
}

func main() {
	var numbers []int

	for {
		var input string
		_, err := fmt.Scanln(&input)
		if err != nil {
			if err.Error() == "EOF" {
				break
			}
			fmt.Println("Ошибка ввода:", err)
			os.Exit(1)
		}

		currentNum, err := strconv.Atoi(strings.TrimSpace(input))
		if err != nil {
			fmt.Println("Ошибка преобразования:", err)
			os.Exit(1)
		}

		numbers = append(numbers, currentNum)

		lower, upper := calculateRange(numbers)

		fmt.Printf("%d %d\n", lower, upper)
	}
}
