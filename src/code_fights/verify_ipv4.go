package main

import "fmt"

func isIPv4Address(inputString string) bool {

	ipSplit := []int{}
	isIP4 := true
	subVal := ""
	dotCount := 0
	totalDots := 0
	numbers := []string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9"}

	if string(inputString[0]) == "." || string(inputString[len(inputString)-1]) == "." {
		return false
	}

	for indx, value := range inputString {

		if string(value) == "." {
			dotCount++
			totalDots++
			if dotCount > 1 {
				return false
			}
			var intOfChar int
			_, _ = fmt.Sscanf(subVal, "%d", &intOfChar)
			ipSplit = append(ipSplit, intOfChar)
			subVal = ""
			continue
		}

		isNotAlpha := false

		for _, numChar := range numbers {
			if string(value) == numChar {
				isNotAlpha = true
				break
			}
		}

		if isNotAlpha == false {
			return isNotAlpha
		}

		dotCount = 0
		subVal = subVal + string(value)

		if indx == len(inputString)-1 {
			var intOfChar int
			_, _ = fmt.Sscanf(subVal, "%d", &intOfChar)
			ipSplit = append(ipSplit, intOfChar)
			subVal = ""
		}
	}

	if totalDots != 3 {
		return false
	}

	for _, ipSub := range ipSplit {
		fmt.Println(ipSplit)
		if ipSub >= 0 && ipSub <= 255 {
			continue
		} else {
			isIP4 = false
		}
	}
	return isIP4
}
