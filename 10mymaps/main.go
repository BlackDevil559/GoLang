package main

import "fmt"

func main() {
	language := make(map[string]string)
	language["Bhumesh"] = "Hindi"
	language["Marathi"] = "Mohit"
	language["Telgu"] = "Neha"
	fmt.Println(language)

	delete(language,"Telgu")
	fmt.Println(language)

	// loops in map
	for key,value:= range(language){
		fmt.Println("key :",key," value:",value)
	}
}