package main

import "fmt"

func main() {
	fmt.Println("welcome to pointers")

	var arr[4] string
	arr[0]="Abhik"
	arr[1]="Abhimanyu-legend!!!!!!!"
	arr[3]="Abinas"
	fmt.Println(arr)
	fmt.Println(len(arr))
	
	var arr_2=[3]int{4,10,22}
	fmt.Println(arr_2)
	fmt.Println(len(arr_2))
}