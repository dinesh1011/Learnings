package main

import "fmt"

//function to sum the digits of number until sum becomes 1 digit

func sumDigits(input int) int {

	sum := 0

	for ; ; {

		if input < 10 && (sum + input) >= 10{
			sum += input
			input = sum
			sum = 0
			continue
		}

		if input < 10 && (sum + input) < 10 {
			return sum + input
		}

		digit := input % 10
		sum += digit
		input = input / 10
	}

}

func main() {

	input := 11232145678999
	fmt.Println(sumDigits(input))

	xx := fmt.Sprintf("%d", 48)
	fmt.Printf("%T\n", xx)


}