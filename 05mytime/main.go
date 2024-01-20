package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("learn about time")

	new_time:=time.Now()
	fmt.Println(new_time)

	fmt.Println(new_time.Format("02-01-2006 15:04:05"))

	create:=time.Date(2024,time.January,20,14,27,23,0,time.Local)
	fmt.Println(create.Format("02-01-2006 15:04:05 Monday"))
}