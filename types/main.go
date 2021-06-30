package main

import "fmt"

func main() {
	var myInt int8
	var myUInt uint8
	myInt = 0
	myUInt = 0
	flag := true
	for myUInt != 0 || flag {
		flag = false
		fmt.Println(myInt, myUInt)
		myInt++
		myUInt++

	}

}
