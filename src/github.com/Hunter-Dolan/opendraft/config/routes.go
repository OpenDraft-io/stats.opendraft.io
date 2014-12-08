package config

import (
	"../app/controllers"
	"github.com/gorilla/mux"
)

func Route(r *mux.Router) {
	//
	// Welcome
	//
	r.HandleFunc("/", controllers.WelcomeShow)

	//
	// Users
	//
	r.HandleFunc("/articles", controllers.HottestArticles)
	r.HandleFunc("/articles/{title}/view", controllers.ViewArticle)
	r.HandleFunc("/articles/{title}/star", controllers.StarArticle)
}
