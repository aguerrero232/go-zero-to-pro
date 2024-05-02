package main

import "fmt"


func isEven(i int) bool {
	return i % 2 == 0
}

func generateNumbersSlice(start int, end int) []int {
	numbers := []int{}
	for i := start; i < end+1; i++ {
		numbers = append(numbers, i)
	}
	return numbers
}


func main () {
	numbers := generateNumbersSlice(0, 10)

	for _, v := range numbers {
		if isEven(v) {
			fmt.Println(v, "is even")
		} else {
			fmt.Println(v, "is odd")
		}
	}
}