package main

import "fmt"

func main() {
	fmt.Println("loops")

	days:=[]string{"monday","tuesday","wednesday","thursday","friday","saturday","sunday"}
	fmt.Println(days)

	//1st way
	/*for i := 0; i < len(days); i++ {
		fmt.Println(days[i])
	}*/
	
	//2nd way
	/*for i:=range days{
		fmt.Println(days[i])
	}*/

	//3rd way

	/*for index,value:=range days{
		fmt.Printf("index is %v with value of %v\n",index,value)
	}*/

	value:=1
	//kind of while loop
	for value<20{

		if(value==10){
			goto value_2
		}
		fmt.Println(value)
		value++
	}

	value_2:
	fmt.Println("This is it")
}