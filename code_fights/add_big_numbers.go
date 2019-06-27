package main

import (
	"fmt"
	"learnings/linked_list"
)

func addTwoHugeNumbers(a *LinkedList.LinkedList, b *LinkedList.LinkedList) *LinkedList.LinkedList {

	num1 := ""
	num2 := ""
	sum := []int{}
	headA := a
	headB := b

	for {
		if headA != nil {
			num1 += fmt.Sprintf("%04d", headA.Value)
			headA = headA.NextElm
		}
		if headB != nil {
			num2 += fmt.Sprintf("%04d", headB.Value)
			headB = headB.NextElm
		}
		if headA == nil {
			num1 = "0000" + num1
		}
		if headB == nil {
			num2 = "0000" + num2
		}
		if headA == nil && headB == nil {
			break
		}

	}

	tempSum := 0
	reminder := 0
	maxLen := len(num1)
	if len(num2) > maxLen {
		maxLen = len(num2)
	}


	for index := maxLen-1; index >= 0; index-- {

		x :=  int(num1[index] - 48)
		y := int(num2[index] - 48)

		tempSum = (x + y + reminder) % 10
		reminder = (x + y + reminder)/10

		sum = append(sum, tempSum)

	}

	fmt.Println(num1, num2, sum)

	output := &LinkedList.LinkedList {0, nil}
	base := output
	sub := 0
	subCount := 0
	isFirst := true

	for count := len(sum) -1; count >=0; count-- {
		sub = (sub * 10) + sum[count]
		subCount++
		if subCount % 4 == 0 || count == 0 {

			if sub == 0 && isFirst == true {
				continue
			}
			output.NextElm = &LinkedList.LinkedList {sub, nil}
			output = output.NextElm
			sub = 0
			isFirst = false
		}
	}

	if base.NextElm == nil {
		return base
	}
	return base.NextElm
}
