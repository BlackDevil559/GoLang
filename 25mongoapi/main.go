package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/BlackDevil559/mongoapi/router"
)

func main() {
	r:=router.Router()
	fmt.Println("Mongo API")
	log.Fatal(http.ListenAndServe(":4000",r))
	fmt.Println("Listening at 4000 Port")
}