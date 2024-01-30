package main

import "fmt"

func main() {
	fmt.Println("Learn about functions")
	//function call
	hola()
	fmt.Println(adder(3,5))

	fmt.Println(pro_adder(3,4,5))
}

func hola(){
	fmt.Println("hola amigos")
}

func adder(value1 int,value2 int)int{
	return value1 +value2;
}

//provide slice as a argument
func pro_adder(values ...int)int{
	total:=0

	for _,val:=range values{
		total+=val
	}

	return total
}