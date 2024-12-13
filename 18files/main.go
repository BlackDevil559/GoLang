package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	content := "This need to go in file."
	file,err:=os.Create("./myfile.txt")
	checkNil(err)
	length,err:=io.WriteString(file,content)
	checkNil(err)
	fmt.Println("The length is:",length)
	file.Close()
	
	reader("./myfile.txt")
}

func reader(filename string){
	data,err:=os.ReadFile(filename)
	checkNil(err)
	fmt.Println(string(data))
}

func checkNil(err error){
	if err!=nil{
		panic(err)
	}
}