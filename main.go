package main

import (
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		text := "hello world"

		w.Write([]byte(text))
	})

	server := new(http.Server)
	server.Handler = mux
	server.Addr = ":5000"

	log.Println("web server is starting at ", server.Addr)

	err := server.ListenAndServe()
	if err != nil {
		log.Fatal(err.Error())
	}
}
