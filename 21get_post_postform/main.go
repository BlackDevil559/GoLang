package main

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
)

func main() {
	fmt.Println("Get Request")
	// getreq("http://localhost:8000/get")
	// postreq("http://localhost:8000/post")
	postformreq("http://localhost:8000/postform")
}

func getreq(s string){
	res,err:=http.Get(s)
	if err!=nil{
		panic(err)
	}
	fmt.Println(res.ContentLength)
	fmt.Println(res.StatusCode)
	defer res.Body.Close()
	data,_:=io.ReadAll(res.Body)
	fmt.Println(string(data))
}

func postreq(s string){
	reqBody:=strings.NewReader(`
	{
		"Name":"Bhumesh",
		"SurName":"Gaur",
		"NickName":"Golu"
	}
	`)
	req,err:=http.Post(s,"application/json",reqBody)
	if err!=nil{
		panic(err)
	}
	fmt.Println(req.StatusCode)
	fmt.Println(req.ContentLength)
	defer req.Body.Close()
	data,_:=io.ReadAll(req.Body)
	fmt.Println(string(data))
}

func postformreq(s string){
	// FormData
	data:=url.Values{}
	data.Add("Name","Bhumesh")
	data.Add("SurName","Gaur")
	data.Add("NickName","Golu")
	req,err:=http.PostForm(s,data)
	if err!=nil{
		panic(err)
	}
	fmt.Println(req.StatusCode)
	fmt.Println(req.Request.ContentLength)
	defer req.Body.Close()
	data1,_:=io.ReadAll(req.Body)
	fmt.Println(string(data1))
}