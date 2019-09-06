package main

import "fmt"

func areSimilar(a []int, b []int) bool {

	if len(a) != len(b) {
		return false
	}

	misPlacedCount := 0

	for counter := range a {

		if a[counter] != b[counter] {
			misPlacedCount++
			bvalCountInA := 0
			bvalCountInB := 0

			for _, avalue := range a {
				if b[counter] == avalue {
					bvalCountInA++
				}
			}

			for _, bvalue := range b {
				if b[counter] == bvalue {
					bvalCountInB++
				}
			}

			if bvalCountInA != bvalCountInB {
				return false
			}
		}
	}

	if misPlacedCount > 2 {
		return false
	}

	return true

}

func palindromeRearranging(inputString string) bool {

	charCount := make(map[string]int)
	isEven := true
	canPalindrome := true

	if len(inputString)%2 > 0 {
		isEven = false
	}

	for _, char := range inputString {
		if _, ok := charCount[string(char)]; ok {
			continue
		}
		for _, element := range inputString {
			if element == char {
				charCount[string(element)]++
			}
		}
	}

	if isEven == true {
		for _, val := range charCount {
			fmt.Println("even", val)
			if (val % 2) != 0 {
				return false
			}
		}
	} else {
		falseCount := 0
		for _, val := range charCount {
			fmt.Println("odd : ", val)
			if (val % 2) != 0 {
				falseCount++
			}
			if falseCount > 1 {
				return false
			}
		}
	}

	return canPalindrome
}
