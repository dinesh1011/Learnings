package main

import (
	"learnings/trees"
	"strings"
)

func binaryTreeSort(inputArray []int, order string){
	initialNode := trees.NewNode(inputArray[0])
	tempBinaryTree := trees.NewBinaryTree(initialNode)

	for counter := 1; counter < len(inputArray); counter++ {
		tempBinaryTree.InsertElement(inputArray[counter])
	}

	if strings.ToLower(order) == "desc" {
		tempBinaryTree.ReadFromRightToLeft()
	}else if strings.ToLower(order) == "asc" {
		tempBinaryTree.ReadFromLeftToRight()
	}else{
		tempBinaryTree.ReadFromLeftToRight()
	}
}

/*
func main(){

	inputArray := []int {135305, 153880, 111255, 21406, 57002, 52054, 103056, 64209,
		41616, 175464, 147703, 72000, 190719, 13351, 137167, 88991, 186086, 99989,
		131110, 99074, 8863, 189535, 80037, 33947, 193112, 405, 188704, 22651, 32, 50654,
		194172, 18810, 93988, 36912, 1680, 152359, 118135, 55270, 13615, 110924, 147293,
		75063, 58826, 51052, 137055, 80180, 93078, 158963, 74302}
	
	binaryTreeSort(inputArray, "desc")

}
*/