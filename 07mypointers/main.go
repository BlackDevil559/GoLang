package main

import "fmt"

func main() {
	fmt.Println("Welcome")
	num:=34;

	// var ptr *int=&num;
	ptr:=&num;
	fmt.Println(ptr);
	fmt.Println(*ptr);
	*ptr+=2
	fmt.Println(ptr);
	fmt.Println(*ptr);
	fmt.Println(num);
}