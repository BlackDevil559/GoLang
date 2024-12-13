package main

import (
	"fmt"
	"io"
	"net/http"
)
const url="https://lco.org"
func main() {
	res, err := http.Get(url)
	if err!=nil{
		panic(err)
	}
	fmt.Printf("Respnse is %T \n",res)
	defer res.Body.Close() // Much necessary to close it.

	data,err:=io.ReadAll(res.Body)
	if err!=nil{
		panic(err)
	}
	content:=string(data)
	fmt.Println(content)
}