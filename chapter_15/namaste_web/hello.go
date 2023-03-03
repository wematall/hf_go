package main

import (
	"log"
	"net/http"
)

func write(writer http.ResponseWriter, message string) {
	_, err := writer.Write([]byte(message))
	if err != nil {
		log.Fatal(err)
	}
}

func hindiHandler(writer http.ResponseWriter, request *http.Request) {
	write(writer, "Namaste web!")
}

func frenchHandler(writer http.ResponseWriter, request *http.Request) {
	write(writer, "Salute web!")
}

func englishHandler(writer http.ResponseWriter, request *http.Request) {
	write(writer, "Hello web!")
}

func main() {
	http.HandleFunc("/namaste", hindiHandler)
	http.HandleFunc("/salute", frenchHandler)
	http.HandleFunc("/hello", englishHandler)
	err := http.ListenAndServe("localhost:8080", nil)
	log.Fatal(err)
}
