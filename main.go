package main

import (
	"log"
	"net/http"

	"github.com/newrelic/go-agent/v3/newrelic"
)

func main() {

	app, err := newrelic.NewApplication(
		newrelic.ConfigAppName("simpleapp"),
		newrelic.ConfigLicense("3bcca23434ab0d41dd9f6a483f31a88dFFFFNRAL"),
		newrelic.ConfigAppLogForwardingEnabled(true),
	)
	mux := http.NewServeMux()
	// mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
	// 	text := "hello world"

	// 	w.Write([]byte(text))
	// })
	mux.HandleFunc(newrelic.WrapHandleFunc(app, "/", func(w http.ResponseWriter, r *http.Request) {
		text := "hello world"

		w.Write([]byte(text))
	}))

	server := new(http.Server)
	server.Handler = mux
	server.Addr = ":5000"

	log.Println("web server is starting at ", server.Addr)

	err = server.ListenAndServe()
	if err != nil {
		log.Fatal(err.Error())
	}
}
