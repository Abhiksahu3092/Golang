package main

import "fmt"

func main() {
	fmt.Println("welcome to pointers")

	var new_number int=45
	var ptr *int
	ptr=&new_number

	fmt.Println(ptr)
	fmt.Println(*ptr)
}