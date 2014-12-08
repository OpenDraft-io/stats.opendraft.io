package controllers

import (
	"fmt"
	"net/http"

	_ "../../config/globals"
	_ "../models"
	_ "github.com/gorilla/mux"
)

func WelcomeShow(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "test")
}
