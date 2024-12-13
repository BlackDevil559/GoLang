package main

import "fmt"

// Use capital letter so that it can be exported
type User struct{
	Name string
	Age int
	Gender string
}
// No inheritance in goLang
func main() {
	fmt.Println("Welcome")
	bhumesh:=User{"Bhumesh",20,"Male"}
	fmt.Println(bhumesh)
	fmt.Printf("Much more detials as %+v",bhumesh)
}
