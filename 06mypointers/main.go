package main

import "fmt"

func main() {
	fmt.Println("welcome to pointers")

	new_number:=45
	var ptr *int
	ptr=&new_number

	fmt.Println(ptr)//address of at which number is stored
	fmt.Println(*ptr)
}