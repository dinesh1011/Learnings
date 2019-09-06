package main

import "fmt"

func variableName(name string) bool {

	isRight := true

	for index, char := range name {

		elmFound := false

		if index == 0 && char >= 48 && char <= 57 {
			return false
		}

		if char >= 65 && char <= 90 {
			fmt.Println("at index 0")
			continue
		}

		if char >= 97 && char <= 122 {
			continue
		}

		if char == 95 {
			continue
		}

		if char >= 48 && char <= 57 {
			continue
		}

		if elmFound == false {
			return elmFound
		}
	}

	return isRight
}

/*
func main() {

	//xx := "Z"
	fmt.Println(2 % 3)

}
*/
