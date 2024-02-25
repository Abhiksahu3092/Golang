package main

import "fmt"

//outside main functions we can use no var type declaration
var new_int="hey everyone"

//here the start of variable with Captial letter has a significance which shows the that the constant is public
const constant_variable=45

func main() {
	var username string="abhik"
	fmt.Println(username)
	fmt.Printf("The variable is of type: %T\n",username)

	var boolean_type bool=false
	fmt.Println(boolean_type)
	fmt.Printf("The variable is of type: %T\n",boolean_type)

	var small_value int=500
	fmt.Println(small_value)
	fmt.Printf("The variable is of type: %T\n",small_value)

	var float_value float64=500.37328462482347324
	fmt.Println(float_value)
	fmt.Printf("The variable is of type: %T\n",float_value)
	
	//here is a comment
	//In go no garbage value is provided if any variable is not initialized
	var not_garbage int
	fmt.Println(not_garbage)
	fmt.Printf("The variable is of type: %T\n",not_garbage)


	//implicit declaration
	var implicit_dec = 45
	fmt.Println(implicit_dec)

	//no var style of declaration
	users:=40.47
	fmt.Println(users)

	fmt.Println(constant_variable)
	fmt.Printf("The variable is of type: %T\n",constant_variable)
	
}