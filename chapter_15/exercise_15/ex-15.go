// Next to each response
// write the URL you'd need to type in
// your browser to generate that response.

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

func d(writer http.ResponseWriter, request *http.Request) {
	write(writer, "z")
}

func e(writer http.ResponseWriter, request *http.Request) {
	write(writer, "x")
}

func f(writer http.ResponseWriter, request *http.Request) {
	write(writer, "y")
}

func main() {
	http.HandleFunc("/a", f) // localhost:4567/a  answer: y
	http.HandleFunc("/b", d) // localhost:4567/b  answer: z
	http.HandleFunc("/c", e) // localhost:4567/c  answer: x

	err := http.ListenAndServe("localhost:4567", nil)
	log.Fatal(err)
}
