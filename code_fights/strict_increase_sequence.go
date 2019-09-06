package main

func almostIncreasingSequenceOptimized(sequence []int) bool {

	isSequence := true
	falseCounter := 0

	for counter := 0; counter < len(sequence)-1; {
		if sequence[counter] >= sequence[counter+1] {
			falseCounter++

			if falseCounter > 1 {
				return false
			}

			if counter+2 > len(sequence)-1 {
				counter++
				continue
			} else if sequence[counter] < sequence[counter+2] {
				counter += 2
				continue
			} else if counter == 0 {
				counter++
				continue
			} else if sequence[counter-1] < sequence[counter+1] {
				counter++
				continue
			} else {
				falseCounter++
			}
		}

		counter++

		if falseCounter > 1 {
			return false
		}
	}
	return isSequence
}

func almostIncreasingSequence(sequence []int) bool {

	isSequence := true

	for ignoreIndex := 0; ignoreIndex < len(sequence); ignoreIndex++ {
		isSequence = true
		newArray := make([]int, len(sequence))
		copy(newArray, sequence)
		newArray = append(newArray[0:ignoreIndex], newArray[ignoreIndex+1:]...)

		for counter := 0; counter < len(newArray)-1; counter++ {
			if newArray[counter] >= newArray[counter+1] {
				isSequence = false
				break
			}
		}

		if isSequence == true {
			return isSequence
		}
	}

	return isSequence
}

/*
func main()  {
	input := []int {3, 6, 5, 8, 10, 20, 15}
	fmt.Println(almostIncreasingSequenceOptimized(input))
}
*/
