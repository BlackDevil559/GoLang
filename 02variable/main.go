package main
import "fmt"
const LoginToken string= "ghatvhibjd" 
// first letter of LoginToken is capital signify it is public keyword
func main(){
	var username string="Bhumesh Gaur"
	fmt.Println(username)
	// Give the type of variable
	fmt.Printf("type of variable is %T \n",username)

	var isloggedin bool=true
	fmt.Println(isloggedin)
	fmt.Printf("type of variable is %T \n",isloggedin)

	var value uint8=16
	fmt.Println(value)
	fmt.Printf("type of variable is %T \n",value)

	// Default value of int is 0.
	var val int
	fmt.Println(val)
	fmt.Printf("type of variable is %T \n",val)

	var value_float float64=165.456789
	fmt.Println(value_float)
	fmt.Printf("type of variable is %T \n",value_float)

	var w="hello"
	fmt.Println(w)
	// It automatically decides the type of w on basis of it's value.
	// w=3 will give an error

	// no var style
	w1:=1220
	fmt.Println(w1)
	// no war style is not allowed outside the function
}
