package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Hii there !!!")

	//takes input from user
	new_input:=bufio.NewReader(os.Stdin)

	//reads the input given by the user
	ok,_:=new_input.ReadString('\n')
	//input is of the form of the string so typecasting is required
	numrating,err:=strconv.ParseFloat(strings.TrimSpace(ok),64)

	if err!=nil {
		fmt.Println(err)
	}else{
		fmt.Println("added 1 to the rating: ",numrating+1)
	}
}