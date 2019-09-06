package main

import "fmt"

func minesweeper(matrix [][]bool) [][]int {

	minesweeperOutput := make([][]int, len(matrix))
	matrixLen := len(matrix[0]) - 1
	matrixListLen := len(matrix) - 1

	for listIndex, matrixlist := range matrix {

		for itemIndex := range matrixlist {

			bombCount := 0
			//fmt.Println("listindex : ", listIndex)
			//fmt.Println("itemindex : ", itemIndex)

			if itemIndex > 0 {
				if matrix[listIndex][itemIndex-1] == true {
					//fmt.Println("in left")
					bombCount++
				}
				if (listIndex < matrixListLen) && (matrix[listIndex+1][itemIndex-1]) == true {
					//fmt.Println("in bottom left corner")
					bombCount++
				}
			}

			if itemIndex < matrixLen {
				if matrix[listIndex][itemIndex+1] == true {
					fmt.Println("in right")
					bombCount++
				}
				if (listIndex > 0) && (matrix[listIndex-1][itemIndex+1] == true) {
					fmt.Println("in top right corner")
					bombCount++
				}
			}

			if listIndex > 0 {
				if matrix[listIndex-1][itemIndex] == true {
					//fmt.Println("in top")
					bombCount++
				}
				if (itemIndex > 0) && (matrix[listIndex-1][itemIndex-1] == true) {
					//fmt.Println("in top left corner")
					bombCount++
				}
			}

			if listIndex < matrixListLen {
				if matrix[listIndex+1][itemIndex] == true {
					//fmt.Println("in bottom")
					//bombCount++
				}
				if (itemIndex < matrixLen) && (matrix[listIndex+1][itemIndex+1] == true) {
					fmt.Println("in bottom right corner")
					bombCount++
				}
			}

			minesweeperOutput[listIndex] = append(minesweeperOutput[listIndex], bombCount)
		}

	}

	return minesweeperOutput

}

func minesweeperOld(matrix [][]bool) [][]int {

	minesweeperOutput := make([][]int, len(matrix))
	matrixLen := len(matrix[0]) - 1
	matrixListLen := len(matrix) - 1

	for listIndex, matrixlist := range matrix {

		for itemIndex := range matrixlist {

			bombCount := 0
			fmt.Println("listindex : ", listIndex)
			fmt.Println("itemindex : ", itemIndex)

			if (itemIndex > 0) && (matrix[listIndex][itemIndex-1] == true) {
				fmt.Println("in left")
				bombCount++
			}

			if (itemIndex < matrixLen) && (matrix[listIndex][itemIndex+1] == true) {
				fmt.Println("in right")
				bombCount++
			}

			if (listIndex > 0) && (matrix[listIndex-1][itemIndex] == true) {
				fmt.Println("in top")
				bombCount++
			}

			if (listIndex < matrixListLen) && (matrix[listIndex+1][itemIndex] == true) {
				fmt.Println("in bottom")
				bombCount++
			}

			//check corners

			if (listIndex > 0) && (itemIndex > 0) && (matrix[listIndex-1][itemIndex-1] == true) {
				fmt.Println("in top left corner")
				bombCount++
			}

			if (listIndex > 0) && (itemIndex < matrixLen) && (matrix[listIndex-1][itemIndex+1] == true) {
				fmt.Println("in top right corner")
				bombCount++
			}

			if (listIndex < matrixListLen) && itemIndex != 0 &&
				(matrix[listIndex+1][itemIndex-1] == true) {
				fmt.Println("in bottom left corner")
				bombCount++
			}

			if (listIndex < matrixListLen) && itemIndex < matrixLen &&
				(matrix[listIndex+1][itemIndex+1] == true) {
				fmt.Println("in bottom right corner")
				bombCount++
			}

			minesweeperOutput[listIndex] = append(minesweeperOutput[listIndex], bombCount)
		}

	}

	return minesweeperOutput

}

/*
func main() {

	matrix := [][]bool{{true, false, false}, {false, true, false}, {false, false, false}}
	fmt.Println(minesweeper(matrix))

}
*/
