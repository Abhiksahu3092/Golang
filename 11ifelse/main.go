package main

import "fmt"

func main() {
	fmt.Println("conditional statements in golang")

	count := 97

	var result string

	if count < 10 {
		result = "regular user"
	}else if count>10{
		result="something coming"
	}else{
		result="exact login count"
	}

	fmt.Println(result)
}
