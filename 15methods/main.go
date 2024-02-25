package main

import "fmt"

func main() {
	fmt.Println("learning about structs")
	
	abhik:=User{"Abhik","abhiksahu2003@gmail.com",true,20}
	fmt.Println(abhik)
	fmt.Printf("%v",abhik.Name)
	fmt.Println()
	
	status(abhik)
	//referencing of these variables is important otherwise a copy of arguement will be executed upon
	fmt.Println(abhik.Age)
	change_age(abhik)
	fmt.Println(abhik.Age)

}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}

//function which takes structure as an argument
func status(myuser User){
	fmt.Println(myuser.Status)
}

func change_age(myuser User){
	myuser.Age=21
	fmt.Println(myuser.Age)
}