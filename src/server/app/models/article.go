package models

import (
	"../../config/globals"
	"math"
	"time"
)

type Article struct {
	Id        int64
	Title     string
	Views     int64
	Stars     int64
	CreatedAt int64
	Score     float64
}

/* Locate Methods */

//
// Finds or Creates Article from the given title
//

func FindOrCreateArticleByTitle(title string) Article {
	article := Article{}
	articles := global.DB.Connection.Model(Article{}).Where("title = ?", title)

	count := 0
	articles.Count(&count)

	if count == 0 {
		article.Title = title
		article.CreatedAt = time.Now().Unix()
		global.DB.Connection.Save(&article)
	} else {
		articles.First(&article)
	}
	return article
}

//
// Returns an array of articles sorted by score then title
//

func FindHottestArticles() []Article {

	var articles []Article

	global.DB.Connection.Order("score desc, title").Limit(10).Find(&articles)

	return articles
}

//
// Updates all the scores of articles older than 1 week but newer than 2 weeks
//

func UpdateAllArticleScores() {
	var articles []Article

	one_week_ago := time.Now().Unix() - 604800.0
	two_weeks_ago := time.Now().Unix() - 1209600.0

	// Score Decay only happens after 1 week or 604800 seconds
	// After 2 weeks or 1209600 scores are so low that they no longer need to be figured

	global.DB.Connection.Where("created_at < ? AND created_at > ?", one_week_ago, two_weeks_ago).Find(&articles)

	for i := range articles {
		articles[i].UpdateScore()
	}
}

/* Struct Methods */

//
// Adds a +1 to the view count
//

func (article Article) View() {
	article.Views = article.Views + 1
	article.UpdateScore()
}

//
// Adds a +1 to the star count
//

func (article Article) Star() {
	article.Stars = article.Stars + 1
	article.UpdateScore()
}

//
// Recalculates the score
//

func (article Article) UpdateScore() {
	power := float64(article.Views + (article.Stars * 4))
	score := math.Log((power + 1))
	age := float64((time.Now().Unix() - article.CreatedAt)) / 604800.0

	// If the article is over 1 week old, score decay will begin
	if age > 1 {
		age = age - 1
		score = score * math.Exp((-8 * age * age))
	}

	article.Score = score
	global.DB.Connection.Save(&article)
}
