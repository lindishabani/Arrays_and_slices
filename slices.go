package main

import "fmt"

var slice[]string

func main() {
	addToSlice("Item 1")
	addToSlice("Item 2")
	showSlice()
}

func addToSlice(a string) {
	fmt.Println("Capacity", cap(slice))
	slice = append(slice, a)
}

func showSlice() {
	if len(slice) > 0 {
		fmt.Println("This is the list:")
		for i := 0; i < len(slice); i++ {
			fmt.Println(slice[i])
		}
	} else {
		fmt.Println("There are no items in the list")
	}
}
