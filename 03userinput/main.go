package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	welcome := "hey welcome here, are you new to golang ?"
	fmt.Println(welcome)
	reader:=bufio.NewReader(os.Stdin)
	fmt.Println("rate our service: ")

	//comma ok or comma error
	input,_:=reader.ReadString('\n')
	fmt.Println("Thanks for rating us ",input)
	fmt.Printf("The type of the input is %T",input)
}