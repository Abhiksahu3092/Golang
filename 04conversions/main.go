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

	new_input:=bufio.NewReader(os.Stdin)

	ok,_:=new_input.ReadString('\n')
	numrating,err:=strconv.ParseFloat(strings.TrimSpace(ok),64)

	if err!=nil {
		fmt.Println(err)
	}else{
		fmt.Println("added 1 to the rating: ",numrating+1)
	}
}