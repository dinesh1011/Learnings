package main

import "fmt"

func reverseInParentheses(inputString string) string {

	while := true
	for ; while == true; {

		while = false
		toReverse := ""
		inReverse := false
		leftPart := ""

		for _, char := range inputString {

			currentElement := string(char)
			//empty the reverse string and
			// turn on flag to add to the string that need to be reversed
			if currentElement == "(" {
				while = true
				inReverse = true
				leftPart += toReverse
				toReverse = ""
			}

			//keep adding to the string to be reversed
			if inReverse == true {
				toReverse += currentElement
				//turn off the flag to stop adding to reverse string
				//reverse the string and add to the left part
				//and empty the string to be reversed
				if currentElement == ")" {
					inReverse = false
					reversed := ""
					for _, char := range toReverse {
						if string(char) == "(" || string(char) == ")" {
							continue
						}
						reversed = string(char) + reversed
					}
					leftPart += reversed
					toReverse = ""
				}
				//use continue to skip adding the same character to left part
				continue
			}
			//if no open or close parenthesis continue adding to left part
			leftPart += currentElement
		}

		inputString = leftPart
	}
	return inputString
}


func main() {

	input := "foo(bar)baz(blim)"

	fmt.Println(reverseInParentheses(input))
	fmt.Println(-(-23))

}
