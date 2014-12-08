package server

import (
	"./config"
	"./config/initializers"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"net/http"
	"os"
)

var debug = true

func Init() {

	initializers.Init(debug)

	r := mux.NewRouter()

	config.Route(r)

	if debug {
		h := handlers.LoggingHandler(os.Stdout, r)
		http.Handle("/", h)
	} else {
		http.Handle("/", r)
	}

	http.ListenAndServe(":8000", nil)
}
