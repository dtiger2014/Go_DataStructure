package main

import(
	"io"
	"net/http"
	"encoding/json"
)


type tttt struct {
	name 	string 	`json:"Name"`
	age 	int32	`json:"Age,omitempty"`
}

func hello(w http.ResponseWriter, r *http.Request) {
	
	v := tttt{
		name:"Hello World",
		age:12,
	}

	str, err := json.Marshal(v)
	if err != nil {
		panic(err.Error())
	}

	io.WriteString(w, string(str))
}

func main() {
	http.HandleFunc("/", hello)
	http.ListenAndServe(":8000", nil)
}