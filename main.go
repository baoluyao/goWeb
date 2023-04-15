package main

import (
	"fmt"
	"log"
	"net/http"
)

func sayHelloName(w http.ResponseWriter, r *http.Request) {
	message := []byte("Hello Word!")
	_, err := w.Write(message)
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	fmt.Println("Hello World!")
	http.HandleFunc("/", sayHelloName)
	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}
