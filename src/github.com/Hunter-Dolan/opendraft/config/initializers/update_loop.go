package initializers

import (
  "./../../app/models"
  "time"
)

//
// Runs the score update method ever 60 seconds
//
// Run this concurrently, otherwise you're going to block everything up

func UpdateScoreLoop() {
  for ;; {
    models.UpdateAllArticleScores()
    time.Sleep(60 *  time.Second)
  }
}