package main

import (
	"fmt"
	"unsafe"
)

func validTime(time string) bool {

	var hrs int
	var secs int

	_, err := fmt.Sscanf(time, "%d:%d", &hrs, &secs)

	fmt.Println(err)

	fmt.Println(hrs, secs)
	appender := ""
	for _, char := range time {

		if char >= 48 && char <= 57 {
			appender += string(char)
			continue
		}
		fmt.Println(appender)
		xx := unsafe.Pointer(&appender)
		fmt.Println(xx)
		hrss := (*[]byte)(xx)
		fmt.Println(hrss)
	}

	fmt.Println(hrs)

return false
}

func main() {

	validTime("12a:44b")

}
