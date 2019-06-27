package main

import (
	"fmt"
	"learnings/linked_list"
)

func reverseNodesInKGroups(l *LinkedList.LinkedList, k int) *LinkedList.LinkedList {

	if k <= 1 || l == nil {
		return l
	}

	length := 1
	head := l
	for head.NextElm != nil {
		length++
		head = head.NextElm
	}

	fmt.Println(length)

	var heads []*LinkedList.LinkedList
	head = l
	navigator := head.NextElm
	head.NextElm = nil

	counter := 1
	indexer := 0

	for {

		if indexer >= (length/k){
			heads = append(heads, navigator)
			break
		}

		if navigator == nil {
			heads = append(heads, head)
			break
		}

		temp := &LinkedList.LinkedList {navigator.Value, nil}
		temp.NextElm = head
		head = temp

		counter++

		navigator = navigator.NextElm


		if counter % k == 0 {
			indexer++
			heads = append(heads, head)
			head = navigator

			if navigator != nil && (length - counter) >= k {
				navigator = navigator.NextElm
			}
			if head != nil && (length - counter) >= k {
				head.NextElm = nil
			}
			counter++

		}

	}

	xx := heads[0]


	for index := 1; index < len(heads); index++ {
		xxNextElm := heads[index]
		for xx != nil && xx.NextElm != nil {
			xx = xx.NextElm
		}
		xx.NextElm = xxNextElm
	}

	return heads[0]

}

func reverseNodesInKGroups1(l *LinkedList.LinkedList, k int) *LinkedList.LinkedList {

	if l == nil || l.NextElm == nil || k == 1 {
		return l
	}
	
	count := 1
	var reversedHead *LinkedList.LinkedList
	head := l
	tail := &LinkedList.LinkedList {0, nil}

loop1:
	for {
		if  (count -1)% k == 0 {
			tail.NextElm = reverseList(head, k)

			if reversedHead == nil {
				reversedHead = tail.NextElm
			}
			tail = head
			
			for index := count; index <= (count-1 + k) && head != nil; index++ {
				head = head.NextElm
				if head == nil {
					if index == (count-1 + k) {
						tail.NextElm = nil
					}
					break loop1
				}
			}
			
			count += k
		}

		count++

	}

	return reversedHead
}

func reverseList(ll *LinkedList.LinkedList, k int) *LinkedList.LinkedList {

	if ll == nil {
		return nil
	}

	revHead := ll
	navigator := ll
	navigator = navigator.NextElm

	for ; k >= 2; k-- {

		if navigator == nil {
			return ll
		}
		temp := &LinkedList.LinkedList {navigator.Value, nil}
		temp.NextElm = revHead
		revHead = temp
		navigator = navigator.NextElm
	}

	return revHead
}

func main () {

	type Ll struct {
		Value int
		next *Ll
	}

	//var myList Ll
	//xx := struct{Value int
						//next Ll} {10, nil}
	myList := &Ll {10, nil}

	xx := myList
	xx.next = &Ll {80, nil}
	//xx.next = myList
	yy := xx.next
	//*yy = Ll {10, nil}
	yy.Value = 30

	//xx.Value = 30
	//myList.Value = 20
	//myList1 := myList
	//*myList1 = Ll{20, nil}
	fmt.Println(yy, xx, xx.next)

	xxy := [3]int{2,4,5}
	yyy := xxy[:]
	//yyy[2] = 6
	fmt.Println(xxy, yyy)

}