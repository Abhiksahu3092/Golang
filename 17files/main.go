package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	//this the method of writing the file
	fmt.Println("we would study about files")
	str := "hello this is Abhik"

	op, err := os.Create("./newfile.txt")

	if err != nil {
		panic(err)
	}

	length, err_2 := io.WriteString(op, str)
	if err_2 != nil {
		panic(err_2)
	}

	fmt.Println("length is: ", length)
	//defer op.Close()

	readfile("./newfile.txt")
}

func readfile(file string) {
	databyte, err := ioutil.ReadFile(file)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(databyte))
}
