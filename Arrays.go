package main

import "fmt"

const capacity int = 5
var list[capacity]string
var length int = 0

func main() {
	addToList("Item 1")
	addToList("Item 2")
	showList()
}

func addToList(a string) {
	if length < capacity {
		list[length] = a
		length++
	} else {
		fmt.Println("Too Many items")
	}
}

func showList() {
	if length > 0 {
		fmt.Println("This is the list:")
		for i := 0; i < length; i++ {
			fmt.Println(list[i])
		}
	} else {
		fmt.Println("There are no items in the list")
	}
}
