package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

// Use the writer to send the greeting to the buffer in our test. Remember fmt.Fprintf is like fmt.Printf but instead takes a Writer to send the string to, whereas fmt.Printf defaults to stdout.
func Greet(writer io.Writer, name string) {
	fmt.Fprintf(writer, "Hello, %s", name)
}

// When you write an HTTP handler, you are given an http.ResponseWriter and the http.Request that was used to make the request. When you implement your server you write your response using the writer.
// http.ResponseWriter also implements io.Writer so this is why we could re-use our Greet function inside our handler.
func MyGreeterHandler(w http.ResponseWriter, r *http.Request) {
	Greet(w, "world")
}

func main() {
	log.Fatal(http.ListenAndServe(":5000", http.HandlerFunc(MyGreeterHandler)))
}