package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Welcome")
	presentTime:=time.Now()
	fmt.Println(presentTime)
	// We have to use it for the particular pattern
	fmt.Println(presentTime.Format("01-02-2006 15:04:05 Monday"))

	createdDate:= time.Date(2024,time.December,10,23,23,0,0,time.UTC)
	fmt.Println(createdDate)
}