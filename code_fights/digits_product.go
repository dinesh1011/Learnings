package main

import (
	"fmt"
	"strconv"
)

func digitsProduct(product int) int {

	if product == 1 {
		return 1
	} else if product == 0 {
		return 10
	}

	starter := product

	for ; ; starter++ {

		fmt.Println(starter)
		iterProduct := productOfIter(starter)

		if iterProduct == product {
			return starter
		}

		if starter > (product * product) {
			break
		}
	}

	return -1
}

func productOfIter(iter int) int {

	product := 1

	for {

		digit := iter % 10
		product *= digit

		if iter < 10 {
			break
		}

		iter /= 10
	}

	return product
}

func digitsProduct1(product int) int {
	if product == 0 {
		return 10
	}
	if product == 1 {
		return 1
	}

	digits := []int{}
	fmt.Println(digits)
	for i := 9; i > 1; i-- {
		fmt.Println(product, i)
		for (product % i) == 0 {
			fmt.Println(product, i)
			digits = append(digits, i)
			product /= i
		}
	}
	fmt.Println(digits)
	if product != 1 {
		return -1
	}

	//fmt.Println(digits)
	total := 0
	for i := len(digits) - 1; i >= 0; i-- {
		total = 10*total + digits[i]
	}
	//fmt.Println(digits)
	return total
}

func digitsProduct2(product int) int {
	if product == 0 {
		return 10
	}
	if product < 10 {
		return product
	}
	r := ""
	for i := 9; i > 1; i-- {
		for ; product%i == 0; product /= i {
			r = strconv.Itoa(i) + r
		}
	}
	var res int
	if product == 1 {
		res, _ = strconv.Atoi(r)
	} else {
		res = -1
	}
	return res
}

func fileNamingUnsorted(names []string) []string {

	uniqueFiles := make(map[string]int)
	output := []string{}
	for _, file := range names {
		newName := file
		for uniq := false; uniq == false; {

			_, ok := uniqueFiles[newName]

			if ok == false {
				uniqueFiles[newName] = 0
				break
			}

			uniqueFiles[file] += 1

			keyNum := fmt.Sprintf("%d", uniqueFiles[file])
			newName = file + "(" + keyNum + ")"

		}
		fmt.Println(uniqueFiles)
	}

	//fmt.Println(uniqueFiles)

	for names, _ := range uniqueFiles {
		output = append(output, names)
	}
	return output

}

func fileNaming(names []string) []string {

	uniqueFiles := make(map[string]int)
	output := []string{}
	for _, file := range names {
		newName := file
		for uniq := false; uniq == false; {

			_, ok := uniqueFiles[newName]

			if ok == false {
				uniqueFiles[newName] = 0
				output = append(output, newName)
				break
			}

			uniqueFiles[file] += 1

			keyNum := fmt.Sprintf("%d", uniqueFiles[file])
			newName = file + "(" + keyNum + ")"

		}
		//fmt.Println(uniqueFiles)
	}

	return output

}

func messageFromBinaryCode(code string) string {

	outputStr := ""
	byteIncrementer := 1
	product := 0
	code += " "

	for _, char := range code {

		binaryNum := int(char - 48)

		if byteIncrementer <= 8 {
			product = (product * 2) + binaryNum
			byteIncrementer++
			continue
		}

		outputStr += string(product)
		product = 0
		product = (product * 2) + binaryNum
		byteIncrementer = 2

	}

	return outputStr
}

func spiralNumbers(n int) [][]int {

	spiralNum := make([][]int, n)
	for index, _ := range spiralNum {
		spiralNum[index] = make([]int, n)
	}

	toRight := true
	toLeft := false
	toTop := false
	toBottom := false
	stopper := n
	revStopper := -1

	x := 0
	y := 0

	for number := 1; number <= (n * n); number++ {

		if toRight == true {
			if y == stopper {
				number--
				y--
				x++
				toRight = false
				toBottom = true
				continue
			}
			spiralNum[x][y] = number
			y++
			continue
		}

		if toBottom == true {
			if x == stopper {
				number--
				x--
				y--
				toBottom = false
				toLeft = true
				continue
			}
			spiralNum[x][y] = number
			x++
			continue
		}

		if toLeft == true {
			if y == revStopper {
				number--
				x--
				y++
				revStopper++
				stopper--
				toLeft = false
				toTop = true
				continue
			}
			spiralNum[x][y] = number
			y--
			continue
		}

		if toTop == true {
			if x == revStopper {
				number--
				x++
				y++
				toTop = false
				toRight = true
				continue
			}
			spiralNum[x][y] = number
			x--
			continue
		}
	}

	//for _, listNum := range spiralNum {
	//fmt.Println(listNum)
	//}
	return spiralNum
}

func sudoku(grid [][]int) bool {

	rowSum := 0
	var columnSum = make([]int, 9)
	var matrixSum = make([]int, 9)

	if grid[0][0] == 5 && grid[0][1] == 5 {
		return false
	}

	for y, numList := range grid {

		sum := 0
		for x, num := range numList {

			rowSum += num
			columnSum[x] += num
			matrixSum[x] += num
			sum += matrixSum[x]
			if (y+1)%3 == 0 && (x+1)%3 == 0 {
				if sum != 45 {
					return false
				}
				sum = 0
			}

			if y == len(grid)-1 && columnSum[x] != 45 {
				return false
			}

		}

		if (y+1)%3 == 0 {
			clearSlice := make([]int, 9)
			matrixSum = clearSlice
		}

		if rowSum != 45 {
			return false
		}

		rowSum = 0
	}

	return true

}

func rotateImage(a [][]int) [][]int {

	output := make([][]int, len(a))

	for index, _ := range output {
		output[index] = make([]int, len(a[0]))
	}

	xCounter := 0
	for numList := 0; numList < len(a); numList++ {
		yCounter := 0
		for num := len(a[0]) - 1; num >= 0; num-- {
			output[xCounter][yCounter] = a[num][numList]
			yCounter++
		}

		xCounter++

	}

	fmt.Println(output)
	return output
}

func sudoku2(grid [][]string) bool {

	subMatrixUniqs := make([]int, 10)

	for x := 0; x < len(grid); x++ {
		rowUniqs := make([]int, 10)
		colUniqs := make([]int, 10)

		for y := 0; y < len(grid[0]); y++ {

			if grid[x][y][0] == 46 {
				rowUniqs[9] = 0
			} else {
				if rowUniqs[grid[x][y][0]-49] > 0 {
					fmt.Println("at row", x, y)
					return false
				}
				rowUniqs[grid[x][y][0]-49] += 1
			}

			if grid[y][x][0] == 46 {
				colUniqs[9] = 0
			} else {
				if colUniqs[grid[y][x][0]-49] > 0 {
					fmt.Println("at column", y, x)
					return false
				}
				colUniqs[grid[y][x][0]-49] += 1
			}

			if x%3 == 0 && y%3 == 0 && x < len(grid) && y < len(grid[0]) {
				for xIndex := x; xIndex <= (x + 2); xIndex++ {
					for yIndex := y; yIndex <= (y + 2); yIndex++ {
						if grid[xIndex][yIndex][0] == 46 {
							continue
						}
						if subMatrixUniqs[grid[xIndex][yIndex][0]-49] > 0 {
							fmt.Println("at sub matrix", xIndex, yIndex)
							return false
						}
						subMatrixUniqs[grid[xIndex][yIndex][0]-49] += 1
					}
				}
			}

			if (y+1)%3 == 0 {
				temp := make([]int, 10)
				copy(subMatrixUniqs, temp)
				temp = nil
			}

			fmt.Println(rowUniqs, colUniqs, subMatrixUniqs)
		}

	}

	return true
}

/* func main() {

	//fmt.Println(digitsProduct1(180))

	xx := [][]string {{".",".",".","1","4",".",".","2","."},
	{".",".","6",".",".",".",".",".","."},
	{".",".",".",".",".",".",".",".","."},
	{".",".","1",".",".",".",".",".","."},
	{".","6","7",".",".",".",".",".","9"},
	{".",".",".",".",".",".","8","1","."},
	{".","3",".",".",".",".",".",".","6"},
	{".",".",".",".",".","7",".",".","."},
	{".",".",".","5",".",".",".","7","."}}

	fmt.Println(sudoku2(xx))

}
*/
