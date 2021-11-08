package main

import (
	"github.com/go-co-op/gocron"
	"github.com/joho/godotenv"
	"github.com/lenistwo/cmd/mailer"
	"github.com/lenistwo/rest"
	"github.com/lenistwo/util"
	"os"
	"strings"
	"time"
)

const (
	keywordDelimiter = ","
	secondsToTimeout = 30
)

var keywords []string

func init() {
	util.CheckError(godotenv.Load("config/.env"))
	rest.New(secondsToTimeout)
	mailer.Setup()
	extractKeywords()
}

func main() {
	forever := make(chan bool)
	createCron()
	<-forever
}

func createCron() {
	s := gocron.NewScheduler(time.Local)
	_, err := s.Cron(os.Getenv("APP_CRON")).Do(checkPromotion)
	util.CheckError(err)
	s.StartAsync()
}

func checkPromotion() {
	promotion := rest.RetrievePromotion()
	if !containsKeyword(promotion.PromotionName) {
		return
	}

	mailer.Send(promotion)
}

func extractKeywords() {
	words := strings.ToLower(os.Getenv("PROMOTION_KEYWORDS"))
	keywords = strings.Split(words, keywordDelimiter)
}

func containsKeyword(productName string) bool {
	for _, k := range keywords {
		if strings.Contains(strings.ToLower(productName), k) {
			return true
		}
	}
	return false
}
