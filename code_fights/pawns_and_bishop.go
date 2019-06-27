package main

func bishopAndPawn(bishop string, pawn string) bool {

	bx := int(bishop[0] - 96)
	by := int(bishop[1] - 48)
	px := int(pawn[0] - 96)
	py := int(pawn[1] - 48)

	x := bx-(by-1)
	y := 1

	if x < 1 {
		y = -(x) + 2
		x = 1
	}

	//find match in increasing diagonal

	for ; y <= 8; {
		if x == px && y == py {
			return true
		}
		x++
		y++
	}

	//find match in decreasing diagonal

	x = 8-((8 - bx) - (by-1))
	y = 1

	if x > 8 {
		y = (x - 8) + 1
		x = 8
	}

	for ; x >=1; {
		if x == px && y == py {
			return true
		}
		x--
		y++
	}

	return false
}
