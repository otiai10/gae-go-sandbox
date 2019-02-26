package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/otiai10/marmoset"
)

func main() {
	router := marmoset.NewRouter()
	router.GET("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hi"))
	})
	router.GET("/hello", func(w http.ResponseWriter, r *http.Request) {
		message := fmt.Sprintf("Hello, %s!", r.FormValue("name"))
		w.Write([]byte(message))
	})

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	fmt.Println("Listening port", port)

	http.ListenAndServe(fmt.Sprintf(":%s", port), router)
}
