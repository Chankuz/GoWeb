package main

import (
	"fmt"
	"log"
	"net/http"
	"webserver/handler"
)

func main() {

	fileServer := http.FileServer(http.Dir("./public"))

	http.Handle("/", fileServer)
	http.HandleFunc("/hello", handler.HelloHanlder)
	http.HandleFunc("/form", handler.FormHanlder)

	fmt.Printf("port running on http://localhost:8080/\n")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	} else {
		fmt.Print("Server is running")
	}

}
