package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("Welcome")
	var fruitlist=[]string{}
	fmt.Printf("The type is %T \n",fruitlist)

	var fruitlist1=[]string{}
	fruitlist1 = append(fruitlist1, "hello")
	fmt.Println(fruitlist1[len(fruitlist1)-1])

	fruitlist1 = append(fruitlist1, "water","maker")
	fmt.Println(fruitlist1)

	// used as slicing tool
	fruitlist1 = append(fruitlist1[1:])
	fmt.Println(fruitlist1)

	t:=make([]int,4)
	t[0]=12
	t[1]=13
	t[2]=9
	t[3]=10
	fmt.Println(t)
	// t[4]=12 error
	t = append(t,6,3,1) // relocate the memory
	fmt.Println(t)
	fmt.Println(sort.IntsAreSorted(t))
	sort.Ints(t)
	fmt.Println(t)
	fmt.Println(sort.IntsAreSorted(t))

	// Delete a keyword for the list
	course:=[]string{"Ruby","ReactJs","JavaScript","GoLang"}
	fmt.Println(course)
	index:=2
	course=append(course[:index],course[index+1:]...)
	fmt.Println(course)
}