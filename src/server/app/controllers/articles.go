package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	//"../../config/globals"
	"../models"
	"github.com/gorilla/mux"
)

func ViewArticle(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(200)

	vars := mux.Vars(r)
	title := vars["title"]

	article := models.FindOrCreateArticleByTitle(title)

	go article.View()

	data := makeRanking(article, -2)
	b, _ := json.Marshal(map[string]RankingArticle{"article_ranking": data})

	fmt.Fprintf(w, string(b))
}

func StarArticle(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(200)

	vars := mux.Vars(r)
	title := vars["title"]

	article := models.FindOrCreateArticleByTitle(title)

	go article.Star()

	data := makeRanking(article, -2)
	b, _ := json.Marshal(map[string]RankingArticle{"article_ranking": data})

	fmt.Fprintf(w, string(b))
}

func HottestArticles(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(200)

	articles := models.FindHottestArticles()

	response := []RankingArticle{}

	for i := range articles {
		data := makeRanking(articles[i], i)
		response = append(response, data)
	}

	b, _ := json.Marshal(map[string][]RankingArticle{"article_rankings": response})

	fmt.Fprint(w, string(b))
}

/*
 * Helper Methods
 */

type RankingArticle struct {
	Id      int64  `json:"id"`
	Title   string `json:"title"`
	Views   int64  `json:"views"`
	Stars   int64  `json:"stars"`
	Score   int64  `json:"score"`
	Ranking int64  `json:"ranking"`
}

func makeRanking(article models.Article, ranking int) RankingArticle {
	data := RankingArticle{}
	data.Id = article.Id
	data.Title = article.Title
	data.Views = article.Views
	data.Stars = article.Stars
	data.Score = int64(article.Score * 1000)
	data.Ranking = int64(ranking + 1)
	return data
}
