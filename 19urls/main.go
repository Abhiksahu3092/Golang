package main

import (
	"fmt"
	"net/url"
)

const new_url string = "https://takeuforward.org/strivers-a2z-dsa-course/strivers-a2z-dsa-course-sheet-2/"

func main() {
	fmt.Println("learning about URLS")
	fmt.Println(new_url)

	//parsing -  The URL parsing functions focus on splitting a URL string into its components, or on combining URL components into a URL string.
	result,_:=url.Parse(new_url)
	fmt.Println(result.Scheme)
	fmt.Println(result.Host)
	fmt.Println(result.Path)
	/*fmt.Println(result.Port())
	fmt.Println(result.RawQuery)*/

	qparams:=result.Query()
	fmt.Printf("%T\n",qparams)

	for _,val:=range qparams{
		fmt.Println("parameter is: ",val)
	}


	partsofUrl:=&url.URL{
		Scheme: "https",
		Host: "climex.dev",
		Path: "/weather",
		RawPath: "user=abhik",
	}

	anotherURL:=partsofUrl.String()
	fmt.Println(anotherURL)
}
