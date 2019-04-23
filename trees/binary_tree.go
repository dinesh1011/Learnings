package trees

import "fmt"

type Node struct {
	value int
	leftNode *Node
	rightNode *Node
}

func NewNode(value int) *Node {
	return &Node{value: value}
}

type binaryTree struct {
	*Node
}

func NewBinaryTree(node *Node) *binaryTree {
	return &binaryTree{Node: node}
}

func (binaryTree *binaryTree) InsertElement(newValue int){
	currentNode := binaryTree.Node
	if currentNode.value ==0 {
		currentNode.value = newValue
		return
	}
	for {
		if currentNode.value >= newValue {
			if currentNode.leftNode == nil {
				currentNode.leftNode = NewNode(newValue)
				break
			}
			currentNode = currentNode.leftNode
			continue
			}
		if currentNode.value < newValue {
			if currentNode.rightNode == nil {
				currentNode.rightNode = NewNode(newValue)
				break
			}
			currentNode = currentNode.rightNode
			continue
		}
	}
}

func (binaryTree *binaryTree) ReadFromLeftToRight(){
	currentNode := binaryTree.Node
	var backTrack []*Node
	backTracked := false

	for {

		if currentNode.leftNode != nil && backTracked == false{
			backTrack = append(backTrack, currentNode)
			currentNode = currentNode.leftNode
			continue
		}else if backTracked == false {
			fmt.Println(currentNode.value)
		}

		if currentNode.rightNode != nil {
			currentNode = currentNode.rightNode
			continue
		}

		if len(backTrack) == 0 {
			break
		}

		currentNode = backTrack[len(backTrack)-1]
		backTracked = true
		backTrack = backTrack[0:(len(backTrack) - 1)]
		fmt.Println(currentNode.value)

		if currentNode.rightNode != nil {
			currentNode = currentNode.rightNode
			backTracked = false
		}else if len(backTrack) == 0 {
			break
		}

	}

}

func (binaryTree *binaryTree) ReadFromRightToLeft(){
	currentNode := binaryTree.Node
	var backTrack []*Node
	backTracked := false

	for {

		if currentNode.rightNode != nil && backTracked == false{
			backTrack = append(backTrack, currentNode)
			currentNode = currentNode.rightNode
			continue
		}else if backTracked == false {
			fmt.Println(currentNode.value)
		}

		if currentNode.leftNode != nil {
			currentNode = currentNode.leftNode
			continue
		}

		if len(backTrack) == 0 {
			break
		}

		currentNode = backTrack[len(backTrack)-1]
		backTracked = true
		backTrack = backTrack[0:(len(backTrack) - 1)]
		fmt.Println(currentNode.value)

		if currentNode.leftNode != nil {
			currentNode = currentNode.leftNode
			backTracked = false
		}else if len(backTrack) == 0 {
			break
		}

	}

}

func (binaryTree *binaryTree) Delete(){
		binaryTree.Node = nil
	return
}

/*
func main() {

	node := NewNode(10)
	bTree := NewBinaryTree(node)

	inputArray := []int{150200, 154099, 19591, 183367, 8, 166508, 83851, 176424}

	for _, elm := range inputArray {
		bTree.InsertElement(elm)
	}

	bTree.ReadFromLeftToRight()

}*/