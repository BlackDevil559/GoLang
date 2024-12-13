package main

import (
	"fmt"
	"math/rand"
)

func main() {
	diceno := rand.Intn(6)+1
	switch diceno{
	case 1:
		fmt.Println("chocie is 1")
	case 2:
		fmt.Println("choice is 2")
	case 3:
		fmt.Println("chocie is 3")
		fallthrough // if we want next also execute
	case 4:
		fmt.Println("choice is 4")
	case 5:
		fmt.Println("chocie is 5")
	case 6:
		fmt.Println("choice is 6")
	}
}
