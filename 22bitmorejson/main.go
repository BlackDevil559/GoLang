package main
import (
	"encoding/json"
	"fmt"
)
type course struct {
	Name     string `json:"coursename"`
	Price    int
	Platform string 
	Password string `json:"-"` // Omit Password from sharing
	Tags     []string `json:"tags,omitempty"`
}
func main() {
	DecodeJson()
}
func EncodeJson() []byte {
	lcoCourses := []course{
		{"ReactJs", 100, "Youtube", "None", []string{"JavaScript", "Node"}},
		{"GoLang", 200, "Youtube", "None", []string{"Python", "CPP"}},
	}

	// Beautify JSON
	finalJson, err := json.MarshalIndent(lcoCourses, "", "\t")
	if err != nil {
		return nil
	}

	return finalJson
}
func DecodeJson(){
	var lcoCourses []course
	jsondata:=EncodeJson()
	fmt.Printf("%s\n",jsondata)
	checkvalid:=json.Valid(jsondata)
	fmt.Println(checkvalid)
	json.Unmarshal(jsondata,&lcoCourses)
	fmt.Printf("%#v\n",lcoCourses)
}