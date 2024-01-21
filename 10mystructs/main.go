package main

import "fmt"

func main() {
	fmt.Println("learning about structs")
	
	abhik:=User{"Abhik","abhiksahu2003@gmail.com",true,20}
	fmt.Println(abhik)
	fmt.Printf("%v",abhik.Name)
	
}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}