package main

import "fmt"

//Tree with node
type Tree struct {
	Value interface{}
	Left  *Tree
	Right *Tree
}

func deleteFromBST1(t *Tree, queries []int) *Tree {
	if t == nil {
		return nil
	}
	for _, item := range queries {

		if t.Value.(int) == item && t.Left == nil && t.Right == nil {
			return nil
		} else if t.Value.(int) == item && t.Left == nil {
			t = t.Right
			continue
		} else if t.Value.(int) == item && t.Left != nil {
			rightMost := 0
			foundRight := false
			findRightMost(t.Left, t.Left, false, &foundRight, &rightMost)
			if foundRight == false {
				t.Value = t.Left.Value
				t.Left = t.Left.Left
			}
			if foundRight == true {
				t.Value = rightMost
			}
			continue
		}

		found := false
		traverseTree(t, t, item, false, &found)
	}
	return t
}

func deleteFromBST(t *Tree, queries []int) *Tree {
	if t == nil {
		return nil
	}
	//var xx = t
	for _, item := range queries {
		if t == nil {
			return nil
		}
		if t.Value.(int) == item && t.Left == nil && t.Right == nil {
			return nil
		}
		deleteItem1(t, item)
	}
	return t
}

func deleteItem1(t *Tree, toRemove int) {

	var backTracker []*Tree
	var parent *Tree
	backTracked := false
	//backTracker = append(backTracker, t)
	right := false

	for {
		if t.Value.(int) == toRemove {
			if t.Left == nil && t.Right != nil {
				t.Value = t.Right.Value
				t.Left = t.Right.Left
				t.Right = t.Right.Right
				break
			} else if t.Left != nil {
				rm := false
				rightMost := 0
				findRightMost1(t.Left, t.Left, false, &rm, &rightMost)
				fmt.Println(rm, rightMost)
				if rm == true {
					t.Value = rightMost
				}
				if rm == false {
					t.Value = t.Left.Value
					t.Left = t.Left.Left
				}
				break
			} else if t.Left == nil && t.Right == nil {
				//fmt.Println("parent", parent, backTracker)
				if right == true {
					parent.Right = nil
				} else {
					parent.Left = nil
				}
			}
			break
		}
		if t.Left != nil && backTracked == false {
			backTracker = append(backTracker, t)
			parent = t
			t = t.Left
			right = false
			//fmt.Println("at left", backTracker[len(backTracker) - 1])
			continue
		}
		if t.Right != nil {
			backTracked = false
			parent = t
			t = t.Right
			right = true
			//fmt.Println("at Right", t)
			continue
		}
		if len(backTracker) == 0 {
			break
		}
		if len(backTracker) != 0 {
			t = backTracker[len(backTracker)-1]
			//parent = t
			//fmt.Println("at backtrack", parent)
			backTracked = true
			backTracker = append(backTracker[0 : len(backTracker)-1])
			continue
		}
	}
}

func findRightMost1(current *Tree, parent *Tree, found bool, rm *bool, retVal *int) {

	if current.Right != nil {
		findRightMost1(current.Right, current, true, rm, retVal)
	}
	if *rm == true {
		return
	}
	if found == true {
		*rm = true
		*retVal = current.Value.(int)
		if current.Left != nil {
			parent.Right = current.Left
		} else {
			parent.Right = nil
		}
	}
}

func findRightMost(current *Tree, parent *Tree, right bool, found *bool, retVal *int) {

	if current.Left != nil {
		findRightMost(current.Left, current, false, found, retVal)
	}

	if current.Right != nil {
		findRightMost(current.Right, current, true, found, retVal)
	}

	if right == true && *found == false {
		*found = true
		*retVal = current.Value.(int)
		if right == true {
			parent.Right = nil
		}
		if right == false {
			parent.Left = nil
		}
	}
}

func deleteItem(t *Tree, toRemove int, found *bool) *Tree {

	if t == nil {
		return nil
	}

	if *found == true {
		return t
	}
	var root *Tree

	if t.Value.(int) == toRemove && t.Left == nil && t.Right == nil {
		*found = true
		root = nil
	} else if t.Value.(int) == toRemove && t.Left == nil {
		*found = true
		root = t.Right
	} else if t.Value.(int) == toRemove && t.Left != nil {
		*found = true
		found := false
		rightMost := 0
		findRightMost(t.Left, t.Left, false, &found, &rightMost)
		if found == true {
			root = &Tree{rightMost, t.Left, t.Right}
		}
		if found == false {
			root = &Tree{t.Left.Value, t.Left.Left, t.Right}
		}
	} else if t.Value.(int) != toRemove {
		root = t
	}

	if *found == true {
		return root
	}

	root.Left = deleteItem(root.Left, toRemove, found)
	root.Right = deleteItem(root.Right, toRemove, found)

	return root
}

func traverseTree(current *Tree, parent *Tree, toRemove int, right bool, found *bool) {

	if *found == true {
		return
	}

	if current.Value.(int) == toRemove {
		*found = true
		if current.Left == nil && right == true {
			parent.Right = current.Right
		}
		if current.Left == nil && right == false {
			parent.Left = current.Right
		}
		if current.Left != nil {
			rightMost := 0
			foundRight := false
			findRightMost(current, current, false, &foundRight, &rightMost)
			if foundRight == true {
				if right == false {
					parent.Left.Value = rightMost
				}
				if right == true {
					parent.Right.Value = rightMost
				}
			}
			if foundRight == false {
				current.Value = current.Left.Value
				current.Left = current.Left.Left
			}
		}
	}

	if current.Left != nil {
		traverseTree(current.Left, current, toRemove, false, found)
	}

	if current.Right != nil {
		traverseTree(current.Right, current, toRemove, true, found)
	}
}

func readTree(tree *Tree, right bool) {
	if tree == nil {
		fmt.Println(nil)
		return
	}
	fmt.Println(tree.Value)
	if tree.Left != nil {
		readTree(tree.Left, false)
	}
	if tree.Right != nil {
		readTree(tree.Right, true)
	}
}

/*
func main() {

	var t Tree
	t.Value = 5
	t.Left = &Tree{2, nil, nil}
	t.Right = &Tree{6, nil, nil}
	t.Left.Left = &Tree{1, nil, nil}
	t.Left.Right = &Tree{3, nil, nil}
	t.Left.Left.Right = &Tree{4, nil, nil}
	t.Left.Right.Right = &Tree{9, nil, nil}
	t.Left.Right.Right.Left = &Tree{10, nil, nil}
	t.Right.Right = &Tree{8, nil, nil}
	t.Right.Right.Left = &Tree{7, nil, nil}
	//t.Right.Right.Left.Right = &Tree{9, nil, nil}
	//t.Right.Right.Right = &Tree{11, nil, nil}

	queries := []int{5}
	//readTree(&t, false)
	xx := deleteFromBST(&t, queries)
	readTree(xx, false)


	str := "fedsxsxsxsx"
	key := 1
	for i, char := range str {
		key += (i+1)*int(char)
	}
	fmt.Println(key)
}
*/
