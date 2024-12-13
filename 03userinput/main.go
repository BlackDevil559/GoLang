package main

import (
	"bufio"
	"fmt"
	"os"
)
func main() {
	var w string = "123"
	fmt.Println(w)

	rating:=bufio.NewReader(os.Stdin)
	fmt.Println("Enter the value of rating ")

	// comma ok \\ comma error syntax
	val,_:=rating.ReadString('\n')
	fmt.Println("Your value is ",val)
	fmt.Printf("your rating type is %T",val)
}