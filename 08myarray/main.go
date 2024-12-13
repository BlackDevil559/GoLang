package main

import "fmt"

func main() {
	var fruitlist [4]string
	fruitlist[0] = "Bhumesh"
	fruitlist[1] = "Gaur"
	fruitlist[2] = "is"
	fruitlist[3] = "fruit"
	fmt.Println("List is",fruitlist)
	fmt.Println("Length is",len(fruitlist))
	fmt.Println("First element is",fruitlist[0])
	// Indexing work fine here like [:2] [:2]
	// [::-1] will not work

	k:=[4]string{"hello","I","am"}
	fmt.Println(k)
}

