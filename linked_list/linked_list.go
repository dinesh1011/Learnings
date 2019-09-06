package linkedlist

import (
	"errors"
	"fmt"
)

// LinkedList - ll object
type LinkedList struct {
	Value   interface{}
	NextElm *LinkedList
}

// AllowedTypes - Restrict your linked list with desired data types
type AllowedTypes []string

func validateType(allowedTypes AllowedTypes, element *LinkedList) (valid bool, err error) {
	var elmType string
	elmType = fmt.Sprintf("%T", element.Value)
	for _, validType := range allowedTypes {
		if validType == elmType {
			return true, nil
		}
	}
	return false, errors.New("invalid type for linked list")
}

// AddItem - Add item to the list
func (lList *LinkedList) AddItem(allowedTypes AllowedTypes, newVal interface{}) {
	newElm := LinkedList{
		Value:   newVal,
		NextElm: nil,
	}

	if validType, _ := validateType(allowedTypes, &newElm); validType != true {
		fmt.Println("invalid type for new element: ", newElm.Value)
		return
	}

	if lList.Value == nil {
		lList.Value = newVal
		lList.NextElm = nil
		return
	}

	head := lList
	for {
		if head.NextElm == nil {
			head.NextElm = &newElm
			return
		} else {
			head = head.NextElm
		}
	}
}

// AddItems - Add multiple items
func (lList *LinkedList) AddItems(allowedTypes AllowedTypes, elements ...interface{}) {
	for _, element := range elements {
		lList.AddItem(allowedTypes, element)
	}
}

// GetItem - Get item by index
func (lList *LinkedList) GetItem(index int) interface{} {
	head := lList
	for elmCount := 1; elmCount < index; elmCount++ {
		head = head.NextElm
	}
	return head.Value
}

// Length - Get length of linked list
func (lList *LinkedList) Length() int {
	head := lList
	length := 0
	for {
		length += 1
		if head.NextElm == nil {
			return length
		} else {
			head = head.NextElm
		}
	}

	return length
}

// RemoveItem - Remove element by index
func (lList *LinkedList) RemoveItem(index int) {

	if lList.Value == nil || index > lList.Length() || index == 0 {
		return
	}

	if index == 1 {
		head := lList
		appender := head.NextElm
		head.Value = appender.Value
		head.NextElm = appender.NextElm
		return
	}

	head := lList
	for elmCount := 2; elmCount <= (index - 1); elmCount++ {
		head = head.NextElm
	}
	appender := head.NextElm.NextElm
	head.NextElm = appender
}

// InsertItem - Insert item to list with index
func (lList *LinkedList) InsertItem(allowedTypes AllowedTypes, newVal interface{}, index int) {

	if index == 0 {
		fmt.Println("invalid index")
		return
	}

	if index > lList.Length() {
		index = lList.Length()
	}

	newElm := LinkedList{
		Value:   newVal,
		NextElm: nil,
	}

	if validType, _ := validateType(allowedTypes, &newElm); validType != true {
		fmt.Println("invalid type for new element: ", newElm.Value)
		return
	}

	head := lList
	for elmCount := 2; elmCount <= index; elmCount++ {
		head = head.NextElm
	}
	newElm.Value = head.Value
	newElm.NextElm = head.NextElm
	head.Value = newVal
	head.NextElm = &newElm

}

// Tests for linked list
func main() {
	var myList LinkedList
	var allowedTypes AllowedTypes
	elements := []interface{}{"new", 10, 23.34, "test", []string{"Dinesh", "Kumar", "Duraisamy"}}
	allowedTypes = append(allowedTypes, "string", "int", "[]string", "float64")
	myList.AddItem(allowedTypes, "Dinesh")
	myList.AddItem(allowedTypes, 100)
	myList.AddItem(allowedTypes, 12)
	myList.AddItems(allowedTypes, elements...)

	for cnt := 0; cnt <= myList.Length(); cnt++ {
		fmt.Println(myList.GetItem(cnt))
	}
	myList.RemoveItem(4)
	myList.InsertItem(allowedTypes, "newTest", 5)
	fmt.Println(myList.Length())

	for cnt := 1; cnt <= myList.Length(); cnt++ {
		fmt.Println(myList.GetItem(cnt))
	}
}
