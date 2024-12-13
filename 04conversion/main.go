package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	b := 7
	fmt.Println(b)
	read:=bufio.NewReader(os.Stdin)
	val,_:=read.ReadString('\n')
	fmt.Println(val)
	intval,err:=strconv.ParseFloat(strings.TrimSpace(val),64)
	if(err!=nil){
		fmt.Println("The error is",err)
	} else{
		fmt.Println(intval+1)
	}
}
