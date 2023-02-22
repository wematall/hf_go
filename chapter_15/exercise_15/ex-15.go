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

func main() {
	http.HandleFunc("/b", d)
	err := http.ListenAndServe("localhost:4567", nil)
	log.Fatal(err)
}
