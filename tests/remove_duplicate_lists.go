package main

import "fmt"

func removeDupList(inputList [][]string) [][]string {

	uniqLists := make(map[string][]string)

	hash := ""

	for _, list := range inputList {
		hash = ""
		for _, element := range list {
			hash += element
		}
		uniqLists[hash] = list
	}

	inputList = nil

	for _, val := range uniqLists {
		inputList = append(inputList, val)
	}

	return inputList
}

func main() {

	input := [][]string {{"1","2","3"}, {"1","2","3","4"},{"1","2","3"},}
	fmt.Println(removeDupList(input))
}