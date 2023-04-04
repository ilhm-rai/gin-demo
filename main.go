package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()

	endpoint := http.HandlerFunc(greet)

	mux.Handle("/", middleware(middleware(endpoint)))

	fmt.Println("Listening to port 8000")

	err := http.ListenAndServe(":8000", mux)

	log.Fatal(err)
}

func greet(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello World!"))
}

func middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println(">")
		next.ServeHTTP(w, r)
	})
}
