package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"

	"github.com/gorilla/mux"
)

type Course struct {
	CourseId    string  `json:"courseid"`
	CourseName  string  `json:coursename`
	CoursePrice int     `json:price`
	Author      *Author `json:author`
}
type Author struct {
	Fullname string `json:"fullname"`
	Website  string `json:"Website"`
}

// Fake DB
var courses []Course

// middleware, helper - file
func (c *Course) Isempty() bool {
	return c.CourseName == ""
}

// Controllers
func serveHome(w http.ResponseWriter, r *http.Request){
	w.Write([]byte("<h1>Welcome to API</h1>"))
}

func getAllCourses(w http.ResponseWriter, r *http.Request){
	fmt.Println("Get all Courses")
	w.Header().Set("Content-Type","application/json")
	json.NewEncoder(w).Encode(courses)
}

func getOneCourse(w http.ResponseWriter, r *http.Request){
	fmt.Println("Get One Course")
	w.Header().Set("Content-Type","application/json")
	params:=mux.Vars(r)
	for _,course:=range(courses){
		if course.CourseId==params["id"]{
			json.NewEncoder(w).Encode(course)
			return 
		}
	}
	json.NewEncoder(w).Encode("No Course is found with given ID")
}

func createOneCourse(w http.ResponseWriter, r *http.Request){
	fmt.Println("Adding a Course")
	w.Header().Set("Content-Type","application/json")
	if(r.Body==nil){
		json.NewEncoder(w).Encode("No data avilable")
		return
	}

	// valid for the get request implementation
	// params:=mux.Vars(r)
	// courses = append(courses, Course{params["id"],params["coursename"],params["price"],params["author"]})
	// json.NewEncoder(w).Encode("Data added sucessfully")

	// for post request
	var course Course
	json.NewDecoder(r.Body).Decode(&course)
	if course.Isempty(){
		json.NewEncoder(w).Encode("Please provide some Valid Data")
		return 
	}
	course.CourseId=string(rand.Intn(100))
	courses = append(courses, course)
	json.NewEncoder(w).Encode(course)
}

func updateOneCourse(w http.ResponseWriter,r *http.Request){
	fmt.Println("Updating the One Course")
	w.Header().Set("Content-Type","application/json")
	params:=mux.Vars(r)
	for index,course:=range(courses){
		if course.CourseId==params["id"]{
			courses=append(courses[:index],courses[index+1:]...)
			var course Course
			json.NewDecoder(r.Body).Decode(&course)
			course.CourseId=params["id"]
			courses = append(courses, course)
			json.NewEncoder(w).Encode(course)
			return 
		}
	}
	json.NewEncoder(w).Encode("ID don't exist")

}

func deleteOneCourse(w http.ResponseWriter,r *http.Request){
	fmt.Println("Deleting the One Course")
	w.Header().Set("Content-Type","application/json")
	params:=mux.Vars(r)
	for index,course:=range(courses){
		if course.CourseId==params["id"]{
			courses=append(courses[:index],courses[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode("Entry Deleted")

}
func main() {
	courses = append(courses, Course{CourseId: "1",CourseName: "ReactJS",
	CoursePrice: 120,Author: &Author{Fullname: "Bhumesh Gaur",Website: "N/A"}})
	r:=mux.NewRouter()
	r.HandleFunc("/",serveHome).Methods("GET")
	r.HandleFunc("/getall",getAllCourses).Methods("GET")
	r.HandleFunc("/getone/{id}",getOneCourse).Methods("GET")
	r.HandleFunc("/deleteone/{id}",deleteOneCourse).Methods("GET")
	r.HandleFunc("/createone",createOneCourse).Methods("POST")
	r.HandleFunc("/updateone/{id}",updateOneCourse).Methods("PUT")

	log.Fatal(http.ListenAndServe(":4000",r))
}