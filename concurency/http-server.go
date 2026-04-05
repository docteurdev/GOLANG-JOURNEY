package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type User struct {
	name string
	age  int
}

func hello(w http.ResponseWriter, req *http.Request) {

	fmt.Fprintf(w, "Hello my backend dev")

}

func getUser(w http.ResponseWriter, req *http.Request) {

	users := []User{
		{name: "djomi dah", age: 20},
		{name: "Foko nounou", age: 60},
		{name: "djomi dah", age: 32},
	}

	u, _ := json.Marshal(&users)

	fmt.Fprint(w, string(u))
}

func main() {

	http.HandleFunc("/hi", hello)

	http.HandleFunc("/users", getUser)

	slcD := []string{"apple", "peach", "pear"}
	slcB, _ := json.Marshal(slcD)
	fmt.Println(string(slcB))

	http.ListenAndServe(":9090", nil)

}
