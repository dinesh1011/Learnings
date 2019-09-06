package main

import "fmt"

func chessKnight(cell string) int {

	x := int(cell[0] - 96)
	y := int(cell[1] - 48)

	counter := 0

	if (x - 2) > 0 {
		if (y + 1) < 9 {
			counter++
		}
		if (y - 1) > 0 {
			counter++
		}
	}
	if (x + 2) < 9 {
		if (y + 1) < 9 {
			counter++
		}
		if (y - 1) > 0 {
			counter++
		}
	}
	if (y - 2) > 0 {
		if (x + 1) < 9 {
			counter++
		}
		if (x - 1) > 0 {
			counter++
		}
	}
	if (y + 2) < 9 {
		if (x + 1) < 9 {
			counter++
		}
		if (x - 1) > 0 {
			counter++
		}
	}

	return counter
}

func main() {

	fmt.Println(chessKnight("g1"))

}
