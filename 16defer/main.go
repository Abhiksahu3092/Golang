package main

import "fmt"

func main() {
	defer fmt.Println("I am defered 1")
	defer fmt.Println("I am defered 2")
	defer fmt.Println("I am defered 3")
	fmt.Println("hey")
	//defer follows a LIFO strategy in defering
	//concept of stacks which works on LIFO will work here
	defering()
}

func defering()  {
	for i := 0; i < 3; i++ {
		defer fmt.Print(i)
	}
	//fmt.Println()
}