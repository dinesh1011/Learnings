package main

import "fmt"

func depositProfit(deposit int, rate int, threshold int) int {

	years := 0
	dep := float64(deposit)
	rat := float64(rate)
	for dep < float64(threshold) {
		years++
		dep += ((dep * rat) / 100)
		fmt.Println(deposit)
	}

	return years

}

/*
func main(){

	fmt.Println(depositProfit(3, 1, 45436454353879))
}
*/
