package main

import "fmt"
func main() {
	days := []string{"sunday", "monday", "tuesday", "thursday", "friday", "saturday"}
	for _,value := range days {
		fmt.Println(value)
	}
	for d:=0;d<len(days);d++{
		fmt.Println(days[d])
		if(d==1){
			goto Lco
		}
	}
	Lco:
	fmt.Println("welcome you buddy")
}