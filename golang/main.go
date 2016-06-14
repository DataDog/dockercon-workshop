package main

import (
	// "encoding/json"
	"fmt"
	"github.com/ChimeraCoder/anaconda"
	"github.com/streadway/amqp"
	"log"
	"net/url"
	"strconv"
	"time"
)

var twitterConsumerKey string = "p2tK5rJTpoF9bDi0c3ekWlACd"
var twitterConsumerSecret string = "nkFXuatVanZjHdzFu7tO5BTYTEL1Q67ZdGUnD23icYvq5V5i38"
var twitterAccessToken string = "2449901-WgodpCBsfDDH7ujv0Q93OwquxquUWLxBLMPhgDlUhI"
var twitterTokenSecret string = "3LMxsFFslbcX74cCZebyy44k0fLQBX8tt5UT5qGPk2ljZ"

var twitterLastRun = time.Now().Add(time.Minute * -5)

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
		panic(fmt.Sprintf("%s: %s", msg, err))
	}
}

// func notLimited(rateLimit int, rateInterval int) bool {
// 	var notlimited = false
// 	var secondsToNextRun = (rateInterval * 60) / rateLimit
// 	if twitterLastRun.Add(time.Second * secondsToNextRun).After(time.Now()) {
// 		notlimited = true
// 	}
// 	return notlimited
// }

func collectTweets(subject string, newestId int64) anaconda.SearchResponse {
	api := anaconda.NewTwitterApi(twitterAccessToken, twitterTokenSecret)
	options := url.Values{}
	options.Set("count", "100")
	options.Set("result_type", "recent")
	options.Set("max_id", strconv.FormatInt(newestId, 10))
	searchResult, _ := api.GetSearch(subject, options)
	return searchResult
}

// func rabbitIsOffline() bool {
// 	result := true
// 	_, err := amqp.Dial("amqp://guest:guest@rabbit:5672")
// 	if err == nil {
// 		result = false
// 	}
// 	return result
// }

func main() {
	// rconnlimiter := time.Tick(time.Second * 5)
	// for rabbitIsOffline() {
	// 	<-rconnlimiter
	// 	log.Println("sleeping till connected")
	// }

	rabbitconn, err := amqp.Dial("amqp://guest:guest@rabbit:5672")

	defer rabbitconn.Close()

	ch, err := rabbitconn.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()

	err = ch.ExchangeDeclare(
		"tweets", // name
		"topic",  // type
		true,     // durable
		false,    // auto-deleted
		false,    // internal
		false,    // no-wait
		nil,      // arguments
	)
	failOnError(err, "Failed to declare an exchange")

	anaconda.SetConsumerKey(twitterConsumerKey)
	anaconda.SetConsumerSecret(twitterConsumerSecret)
	maxId := int64(9223372036854775807)
	limiter := time.Tick(time.Second * 60 * 15 / 180)
	for {
		<-limiter
		tweets := collectTweets("dockercon OR docker OR datadog", maxId)

		for _, tweet := range tweets.Statuses {
			body := []byte(fmt.Sprintf("%v - %v - %v - %v", strconv.FormatInt(tweet.Id, 10), tweet.CreatedAt, tweet.User.ScreenName, tweet.Text))
			// fmt.Println(tweet.Text)
			maxId = tweet.Id
			err = ch.Publish(
				"tweets",   // exchange
				"logstash", // routing key
				false,      // mandatory
				false,      // immediate
				amqp.Publishing{
					ContentType: "text/plain",
					Body:        body,
				})
			failOnError(err, "Failed to publish a message")

		}

	}
}
