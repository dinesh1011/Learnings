package main

import "fmt"

func boxBlur(image [][]int) [][]int {

	xOutLen := len(image) - 2
	yOutLen := len(image[0]) - 2
	averagePixes := make([][]int, xOutLen)

	for index, _ := range averagePixes {
		averagePixes[index] = make([]int, yOutLen)
	}

	minImageSize := 3
	subArrayXPosition := 0
	subArrayYPosition := 0

	for xNavigator :=0; xNavigator < xOutLen; xNavigator++ {
		//fmt.Println("subX:", subArrayXPosition)
		//fmt.Println("subY:", subArrayYPosition)
		//fmt.Println("xoutlen : ", xOutLen)
		subArrayYPosition = 0
		for yNavigator := 0; yNavigator < yOutLen; yNavigator++ {
			//fmt.Println("yNavigator : ",  yNavigator)
			//fmt.Println("yOutLen : ",  yOutLen)
			//fmt.Println("subArrayYPosition : ",  subArrayYPosition)
			pixelsSum := 0
			x := subArrayXPosition

			for xCounter := 0; xCounter < minImageSize; xCounter++ {

				y := subArrayYPosition

				for yCounter := 0; yCounter < minImageSize; yCounter++ {
					//fmt.Println("image pos :" , x, y)
					pixelsSum += image[x][y]
					y++
				}

				x++

			}

			average := (pixelsSum / 9)
			averagePixes[subArrayXPosition][subArrayYPosition] = average
			subArrayYPosition++

		}

		subArrayXPosition++
	}

	fmt.Println(averagePixes)
	return averagePixes
}

func main(){

	image := [][]int{{36,0,18,9},
	{27,54,9,0},
	{81,63,72,45}}
	boxBlur(image)
}

