package main

import (
	"fmt"
	"net/http"
)

func formHandler(res http.ResponseWriter, req *http.Request) {
	if err := req.ParseForm(); err != nil {
		fmt.Fprintf(res, "parseForm() err: %v", err)
		return
	}

	fmt.Println("Post request is successful")

	name := req.FormValue("name")
	address := req.FormValue("address")

	fmt.Fprintf(res, "the name is: %s\n", name)
	fmt.Fprintf(res, "the address is: %s\n", address)
}

func main() {
	fileServer := http.FileServer(http.Dir("./static"))
	http.Handle("/", fileServer)

	http.HandleFunc("/form", formHandler)

	http.ListenAndServe(":9090", nil)
	fmt.Print("the server is running on http://localhost:9090")
}
