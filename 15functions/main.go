package main

import "fmt"

// Nested functions are not allowed
func greet() {
	fmt.Println("Hello from gaur family")
}
func adder(val1 int,val2 int) int{
	return val1+val2
}
func proAdder(val... int) int{
	total:=0
	for _,val:=range(val){
		total+=val
	}
	return total
}
func gen()(string,int){
	return "bhumesh",7
}
func main() {
	greet()
	a:=adder(2,3)
	b:=proAdder(3,4,5,6,7)
	fmt.Println(a,b)
	fmt.Println(gen())
}