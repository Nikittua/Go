package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func sayhello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Привет!")
}

func currtime(w http.ResponseWriter, r *http.Request) {
        dt :=time. Now() 

	fmt.Fprintf(w, dt.String())
}


func main() {
//	fmt.Println("Hello, World!")
	http.HandleFunc("/", sayhello)// Устанавливаем роутер
	http.HandleFunc("/time", currtime)
	err := http.ListenAndServe(":8080", nil) // устанавливаем порт веб-сервера
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
