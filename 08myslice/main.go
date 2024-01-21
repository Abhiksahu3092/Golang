package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("slices")

	var new_arr=[]int{1,4,23,30}
	fmt.Println(new_arr)

	new_arr=append(new_arr,24,25,26)
	fmt.Println(new_arr)
	
	//new_arr=append(new_arr[1:4])
	fmt.Println(new_arr)
	
	//new_arr=append(new_arr[:4])
	fmt.Println(new_arr)

	//makes a slice 
	highscores:=make([]int, 4)

	highscores[0]=24
	highscores[1]=25
	highscores[2]=26
	highscores[3]=27


	highscores=append(highscores, 45,434,67)

	sort.Ints(highscores)

	fmt.Println(highscores)

	fmt.Println(sort.IntsAreSorted(highscores))
	
	var courses=[]string{"c++","javascript","python","golang"}
	fmt.Println(courses)
	
	//we would study about this later on
	index:=2
	courses=append(courses[:index],courses[index+1:]...)
	fmt.Println(courses)
}