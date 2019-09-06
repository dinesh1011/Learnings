package main

import "fmt"

func deleteDigit(n int) int {
	mul, max := 1, 0
	var newNumber int
	for mul < n {
		newNumber = ((n / (mul * 10)) * mul) + (n % mul)
		fmt.Println(newNumber)
		if max < newNumber {
			max = newNumber
		}
		mul *= 10
	}
	return max
}

/* func main() {

	fmt.Println(deleteDigit(232312))

} */
