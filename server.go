package main

import (
	"fmt"
	"log"
	"net/http"
)

func helloController(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/hello" {
		http.Error(w, "404 Page not found", http.StatusNotFound)
		return
	}
	if r.Method != "GET" {
		http.Error(w, "Method not supported", http.StatusNotFound)
		return
	}
	fmt.Fprintf(w, "Hello world !")
}

func formController(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Method not supported", http.StatusNotFound)
		return
	}
	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "ParseForm err %v", err)
		return 
	}
	fmt.Fprintf(w, "Post request success !\n")
	name := r.FormValue("name")
	number := r.FormValue("number")
	fmt.Fprintf(w, "User name: %s\n", name)
	fmt.Fprintf(w, "User number: %s\n", number)
}

func main() {
	fileServer := http.FileServer(http.Dir("./static"))
	http.Handle("/", fileServer)
	http.HandleFunc("/hello", helloController)
	http.HandleFunc("/form", formController)

	fmt.Printf("Listening http server at port 8080 ...\n")
	if err := http.ListenAndServe(":8080", nil); err != nil{
		log.Fatal(err)
		return
	}
}