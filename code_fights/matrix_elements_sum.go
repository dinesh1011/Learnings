package main

func matrixElementsSum(matrix [][]int) int {
	matrixSum := 0
	for rowIndex, row := range matrix {
		for elemIndex, element := range row {
			// check element is not above 0
			if rowIndex == 0 {
				matrixSum += element
				continue
			}

			add := true
			for counter := rowIndex - 1; counter >= 0; counter-- {
				if matrix[counter][elemIndex] == 0 {
					add = false
					break
				}
			}

			if add == true {
				matrixSum += element
			}
		}
	}

	return matrixSum

}

func matrixElementsSum1(matrix [][]int) int {

	sum := 0
	for x := 0; x < len(matrix[0]); x++ {
		for y := 0; y < len(matrix) && matrix[y][x] != 0; y++ {
			sum += matrix[y][x]
		}
	}

	return sum
}

/*
func main(){

	input := [][]int{{0, 1, 1, 0},{1, 5, 0, 1},{2, 1, 3, 10}}
	fmt.Println(matrixElementsSum(input))

	fmt.Println(len("xxz  "))

	}

*/
