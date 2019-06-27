package main

import "fmt"

func strstr(s string, x string)int {

	str := s
	subStr := x
	if len(str) < len(subStr){
		return -1
	}
	if len(str) == len(subStr){
		if str == subStr {
			return 0
		}
		return -1
	}

	currentSubChar := 0
	nextSubChar := 0
	currentSubCharCount := 0
	currentSubIndex := 0
	subCharMatch := false
	matchedIndex := -1
	preIndex := 0
	start := true

	currentSubChar, nextSubChar, currentSubCharCount, currentSubIndex, subCharMatch = setSubstrMatchChars(subStr, currentSubIndex)

	if len(subStr) == 1{
		currentSubChar = int(subStr[0])
	}

	for index := 0; index < len(str) - 1; index++ {
		if int(str[index]) == currentSubChar && nextSubChar == -1 {
			subCharCount := currentSubCharCount
			subCharMatch = true
			for sI := index; sI < len(str) && currentSubCharCount >= 0; sI++ {
				if int(str[sI]) !=currentSubChar {
					subCharMatch = false
					currentSubIndex = 0
					break
				}
				currentSubCharCount--
			}
			if currentSubCharCount > 0 {
				subCharMatch = false
				currentSubIndex = 0
				index += currentSubCharCount
			}else{
				index += subCharCount
			}

			if subCharMatch == true {
				return index + 1 - (len(subStr))
			}
			if  currentSubIndex < len(subStr) - 1{
				currentSubChar, nextSubChar, currentSubCharCount, currentSubIndex, subCharMatch = setSubstrMatchChars(subStr, currentSubIndex)
			}
		}else if int(str[index]) == currentSubChar && int(str[index + 1]) == nextSubChar {
			subCharMatch = true
			SubCharCount := currentSubCharCount
			for sIndex := index - 1; sIndex >= 0 && currentSubCharCount > 0 ; sIndex-- {
				if int(str[sIndex]) !=currentSubChar {
					subCharMatch = false
					currentSubIndex = 0
					//index = sIndex
					//start = true
					break
				}
				currentSubCharCount--
			}

			if start == false && subCharMatch == true && (index - preIndex) != SubCharCount +1 {
				preIndex = index
				start = true
				currentSubIndex = 0
				subCharMatch = false
			}else {
				preIndex = index
				start = false
			}
			if currentSubCharCount > 0 {
				subCharMatch = false
				currentSubIndex = 0
			}
			if  currentSubIndex < len(subStr) - 1{
				currentSubChar, nextSubChar, currentSubCharCount, currentSubIndex, subCharMatch = setSubstrMatchChars(subStr, currentSubIndex)
			}
		}

		if start == false && (index - preIndex) > currentSubCharCount {
			index = preIndex + 1
			preIndex = index
			start = true
			currentSubIndex = 0
			subCharMatch = false
			currentSubChar, nextSubChar, currentSubCharCount, currentSubIndex, subCharMatch = setSubstrMatchChars(subStr, currentSubIndex)
		}

		if currentSubIndex >= len(subStr) - 1 && subCharMatch == true{
			matchedIndex = index
			break
		}

	}

	if matchedIndex < 0 {
		return matchedIndex
	}

	return matchedIndex - (len(subStr) - 2)
}

func setSubstrMatchChars(subStr string, currentSubIndex int)(int, int,int, int, bool){

	currentSubChar := -1
	nextSubChar := -1
	currentSubCharCount := 0

	for subIndex := currentSubIndex; subIndex < len(subStr) - 1; subIndex++ {
		currentSubIndex = subIndex + 1
		currentSubChar = int(subStr[subIndex])
		if subStr[subIndex] != subStr[subIndex + 1] {
			nextSubChar = int(subStr[subIndex + 1])
			break
		}
		currentSubCharCount++
	}

	//fmt.Println("at set", currentSubChar, nextSubChar, currentSubCharCount, currentSubIndex)
	return currentSubChar, nextSubChar, currentSubCharCount, currentSubIndex, false
}

func main(){
	input := "aaaaaaaaaaaaaaaaxaaaaaaaaaaaaaaaaaaaaaaxaaxaabaaaaaaaaaaaaaaaaxaaaaaaaaaaaaaaaaaaaaaaaaaaaab"
	sub := 	"aaxaaaaaaaaaaaaaaaaaxaaaaaaaaaab"
	fmt.Println(strstr(input, sub))
}