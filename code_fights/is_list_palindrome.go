package main

import (
	"fmt"
	"learnings/linked_list"
	"time"
)

func isListPalindrome(l *LinkedList.LinkedList) bool {

	head := l
	if head == nil {
		return true
	}
	actual := []string{}
	val := ""

	for ; ; {

		val = fmt.Sprintf("%d", head.Value.(int))
		actual = append(actual, val)
		if head.NextElm == nil {
			break
		}
		head = head.NextElm
	}

	decrementer := len(actual)-1

	for _, str := range actual {
		if str != actual[decrementer] {
			return false
		}
		decrementer--
	}

	return true

}

func isListPalindromeRecur(l *LinkedList.LinkedList) bool {

	head := l
	check, _ := isPalindrome (head, l)
	return check

}

func isPalindrome(head *LinkedList.LinkedList, base *LinkedList.LinkedList) (bool,  *LinkedList.LinkedList) {

	if head == nil {
		return true, base
	}

	check, base := isPalindrome(head.NextElm, base)

	if check == false {
		return false, base
	}

	if head.Value != base.Value {
		return false, base
	}

	base = base.NextElm

	return true, base

}

func reverseLinkList(l * LinkedList.LinkedList)  LinkedList.LinkedList {

	head := l
	navigator := l.NextElm
	head.NextElm = nil

	for   {
		tempVal := LinkedList.LinkedList{navigator.Value, nil}
		tempVal.NextElm = head
		head = &tempVal
		if navigator.NextElm == nil {
			break
		}
		navigator = navigator.NextElm
	}

return *head

}


func rearrangeLastN(l *LinkedList.LinkedList, n int) *LinkedList.LinkedList {

	if l == nil || n < 1 {
		return l
	}

	var midHead *LinkedList.LinkedList
	var midTail *LinkedList.LinkedList

	head := l
	tail := l
	lenHead := l
	length := 1
	counter := 1

	for lenHead.NextElm != nil {
		length++
		lenHead = lenHead.NextElm
	}
	if n >= length {
		return l
	}
	for head.NextElm != nil {

		if (length - counter) == n-1 {
			//fmt.Println("a", counter)
			midHead = head
		}
		if (length - counter) == n {
			//fmt.Println("b", counter)
			midTail = head
		}

		counter++
		head = head.NextElm
	}

	if (length - counter) == n-1 {
		midHead = head
	}
	tail = head
	tail.NextElm = l
	midTail.NextElm = nil

	return midHead

}


func main(){

	var myList LinkedList.LinkedList
	var allowedTypes LinkedList.AllowedTypes
	allowedTypes = append(allowedTypes, "int")
	elements := []interface{} {1, 2, 3, 4, 5, 6, 7}
	
	myList.AddItems(allowedTypes, elements...)

	fmt.Println("starting check", time.Now())
	xx := rearrangeLastN(&myList, 1)
	for {
		fmt.Println(xx)
		xx = xx.NextElm
		if xx == nil {
			break
		}
	}
	fmt.Println("End check" , time.Now())
}