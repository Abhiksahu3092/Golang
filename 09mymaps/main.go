package main

import "fmt"

func main() {
	fmt.Println("Maps in golang")
	
	languages:=make(map[string]string)

	languages["JS"]="javascript"
	languages["py"]="python"
	languages["bs"]="bootstrap"


	fmt.Println(languages["JS"])
	
	delete(languages,"JS")
	fmt.Println(languages)

	//for loops

	for index, value := range languages {
		fmt.Printf("%v for index %v\n",value,index)
	}
	
}