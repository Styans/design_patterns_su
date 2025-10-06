package kiss

import "fmt"

func ProcessNumbers(numbers []int) {
	if len(numbers) == 0 {
		return
	}
	for _, number := range numbers {
		if number > 0 {
			fmt.Println(number)
		}
	}
}

func PrintPositiveNumbers(numbers []int) {
	for _, n := range numbers {
		if n > 0 {
			fmt.Println(n)
		}
	}
}

func Divide(a, b int) int {
	if b == 0 {
		return 0
	}
	return a / b
}
