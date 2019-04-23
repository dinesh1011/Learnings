package main

import (
	"errors"
	"fmt"
)

type LinkedList struct {
	value interface{}
	nextElm * LinkedList
}

// Restrict your linked list with desired data types
type AllowedTypes []string

func validateType(allowedTypes AllowedTypes, element *LinkedList) (valid bool, err error) {
	var elmType string
	elmType = fmt.Sprintf("%T", element.value)
	for _, validType := range allowedTypes {
		if validType == elmType{
			return true, nil
		}
	}
	return false, errors.New("invalid type for linked list")
}

// Add item to the list
func (lList *LinkedList) AddItem(allowedTypes AllowedTypes, newVal interface{}) {
	newElm := LinkedList{
		value: newVal,
		nextElm: nil,
	}

	if validType, _ := validateType(allowedTypes, &newElm); validType != true {
		fmt.Println("invalid type for new element: ", newElm.value)
		return
	}

	if lList.value == nil{
		lList.value =newVal
		lList.nextElm = nil
		return
	}

	head := lList
	for {
		if head.nextElm == nil {
			head.nextElm = &newElm
			return
		}else {
			head = head.nextElm
		}
	}
}

// Add multiple items
func (lList *LinkedList) AddItems(allowedTypes AllowedTypes,  elements ...interface{}){
	for _, element := range elements{
		lList.AddItem(allowedTypes, element)
	}
}

// Get item by index
func (lList *LinkedList) GetItem(index int) interface{} {
	head := lList
	for elmCount := 1; elmCount < index; elmCount++ {
			head = head.nextElm
	}
	return head.value
}

// Get length of linked list
func (lList *LinkedList) Length() int {
	head := lList
	length := 0
	for {
		length += 1
		if head.nextElm == nil {
			return length
		}else {
			head = head.nextElm
		}
	}

	return length
}

// Remove element by index
func (lList *LinkedList) RemoveItem(index int) {
	if lList.value == nil || index > lList.Length() || index == 0{
		return
	}

	if index == 1 {
		head := lList
		appender := head.nextElm
		head.value = appender.value
		head.nextElm = appender.nextElm
		return
	}

	head := lList
	for elmCount :=2; elmCount <= (index -1); elmCount ++{
			head = head.nextElm
		}
	appender := head.nextElm.nextElm
	head.nextElm = appender
	}

// Insert item to list with index
func (lList *LinkedList) InsertItem(allowedTypes AllowedTypes, newVal interface{}, index int){

	if index == 0 {
		fmt.Println("invalid index")
		return
	}

	if index > lList.Length(){
		index = lList.Length()
	}

	newElm := LinkedList{
		value: newVal,
		nextElm: nil,
	}

	if validType, _ := validateType(allowedTypes, &newElm); validType != true{
		fmt.Println("invalid type for new element: ", newElm.value)
		return
	}

	head := lList
	for elmCount := 2; elmCount <= index; elmCount++ {
			head = head.nextElm
	}
	newElm.value = head.value
	newElm.nextElm = head.nextElm
	head.value = newVal
	head.nextElm = &newElm

}


// Tests for linked list
func main(){
	var myList LinkedList
	var allowedTypes AllowedTypes
	elements := []interface{}{"new", 10, 23.34, "test", []string {"Dinesh", "Kumar", "Duraisamy"}}
	allowedTypes = append(allowedTypes, "string", "int", "[]string", "float64")
	myList.AddItem(allowedTypes, "Dinesh")
	myList.AddItem(allowedTypes, 100)
	myList.AddItem(allowedTypes, 12)
	myList.AddItems(allowedTypes, elements...)

	for cnt:=0 ; cnt <= myList.Length() ; cnt++{
		fmt.Println(myList.GetItem(cnt))
	}
	myList.RemoveItem(4)
	myList.InsertItem(allowedTypes, "newTest", 5)
	fmt.Println(myList.Length())

	for cnt:=1 ; cnt <= myList.Length() ; cnt++{
		fmt.Println(myList.GetItem(cnt))
	}
}