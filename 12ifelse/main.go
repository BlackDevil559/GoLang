package main

import "fmt"

func main() {
	fmt.Println("If else condition")
	val:=12
	if(val==2){
		fmt.Println("val is 2")
	}else if(val==12) {
		fmt.Println("val is 12")
	} else{
		fmt.Println(val)
	}

	if(10>2 && 10>3){
		fmt.Println("Huarrha")
	}else{
		fmt.Println("so sad")
	}

}
