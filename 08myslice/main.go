package main

import "fmt"

func main() {
	fmt.Println("slices")

	var new_arr=[]int{1,4,23,30}
	fmt.Println(new_arr)

	new_arr=append(new_arr,24,25,26)
	fmt.Println(new_arr)
	
	new_arr=append(new_arr[1:4])
	fmt.Println(new_arr)
	
	new_arr=append(new_arr[:4])
	fmt.Println(new_arr)
}