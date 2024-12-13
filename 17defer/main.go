package main

import "fmt"

// defer statement can be assumed as executed in the last.
// It is Last in Fast out
func main() {
	defer fmt.Println("Hello")
	defer fmt.Println("One")
	defer fmt.Println("Two")
	fmt.Println("World")
}
