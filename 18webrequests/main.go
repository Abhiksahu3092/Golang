package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

const url = ("https://climex.vercel.app/")

func main() {
	fmt.Println("handling web requests")

	response, err := http.Get(url)

	if err != nil {
		panic(err)
	}

	fmt.Printf("%T", response)

	defer response.Body.Close()
	//callers responsibility to close the connection

	db, err := ioutil.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(db))
}
