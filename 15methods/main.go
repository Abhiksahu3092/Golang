package main

import "fmt"

func main() {
	fmt.Println("learning about structs")
	
	abhik:=User{"Abhik","abhiksahu2003@gmail.com",true,20}
	fmt.Println(abhik)
	fmt.Printf("%v",abhik.Name)
	fmt.Println()
	
	abhik.status()
	//referencing of these variables is important otherwise a copy of arguement will be executed upon
	fmt.Println(abhik.Age)
	abhik.change_age()
	fmt.Println(abhik.Age)

}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}

//function which takes structure as an argument
func (myuser User)status(){
	fmt.Println(myuser.Status)
}

func (myuser User)change_age(){
	myuser.Age=21
	fmt.Println(myuser.Age)
}