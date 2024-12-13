// Functions in structs are Methods
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
	fmt.Printf("Much more detials as %+v\n",bhumesh)
	bhumesh.Getage()
	fmt.Println(bhumesh)
}

func (u User) Getage(){
	// Passes a copy,original value remain unchanged
	u.Name="QWERTY"
	fmt.Println("The age is",u.Age,u.Name)	
}