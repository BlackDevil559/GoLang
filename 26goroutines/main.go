package main

import (
	"fmt"
	"net/http"
	"sync"
)
var wg sync.WaitGroup
func main() {
	endPoints:=[]string{"https://youtube.com",
						"https://google.com",
						"https://instagram.com"}
	for _,val:=range(endPoints){
		go getStatusCode(val)
		wg.Add(1)
	}
	wg.Wait()
}

func getStatusCode(endPoint string) {
	defer wg.Done()
	res,err:=http.Get(endPoint)
	if err!=nil{
		panic(err)
	}
	fmt.Printf("%d status code for %s \n",res.StatusCode,endPoint)
}

