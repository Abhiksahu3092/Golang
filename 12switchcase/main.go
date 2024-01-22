package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println("learning about switch cases")

	rand.Seed(time.Now().UnixNano())
	dicenumber:=rand.Intn(6)+1

	fmt.Println("dice number is: ",dicenumber)

	switch dicenumber {
	case 1:
		fmt.Println("you can jump 1 place")
		//fallthrough
		//fallthrough is something like a continue statement in switch cases 
	case 2:
		fmt.Println("you can jump 2 places")
	case 3:
		fmt.Println("you can jump 3 places")
	case 4:
		fmt.Println("you can jump 4 places")
	case 5:
		fmt.Println("you can jump 5 places")
	case 6:
		fmt.Println("you can jump 6 places or you can open up")
	default:
		fmt.Println("invalid")
	}
}