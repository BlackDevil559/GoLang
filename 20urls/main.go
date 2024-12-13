package main

import (
	"fmt"
	"net/url"
)

const murl = "https://lco.dev:3000/learn?coursename=reactjs&age=12"

func main() {
	fmt.Println(murl)
	result,_:=url.Parse(murl)
	fmt.Println(result.Scheme)
	fmt.Println(result.Path)
	fmt.Println(result.Host)
	fmt.Println(result.RawQuery)
	fmt.Println(result.Port())

	q:=result.Query()
	fmt.Printf("%T \n",q)
	fmt.Println(q["coursename"])
	for key,value := range(q){
		fmt.Println(key,value)
	}

	part:=	&url.URL{
		Scheme: "https",
		Host:"lco.dev",
		Path: "tutcss",
		RawQuery: "user=hitesh",
	}
	anotherURL:=part.String()
	fmt.Println(anotherURL)
}	