package main

import (
	"fmt"
	"github.com/ChimeraCoder/anaconda"
	"github.com/DataDog/datadog-go/statsd"
	"github.com/streadway/amqp"
	"log"
	"net/url"
	"os"
	"strconv"
	"time"
  "net/http"
)

var twitterConsumerKey string = os.Getenv("twitterConsumerKey")
var twitterConsumerSecret string = os.Getenv("twitterConsumerSecret")
var twitterAccessToken string = os.Getenv("twitterAccessToken")
var twitterTokenSecret string = os.Getenv("twitterTokenSecret")
var coolTagForApp string = "app:twittergolang"
var twitterLastRun = time.Now().Add(time.Minute * -5)
var aggregationKey = "GoLangApp"

var host string
var apiRateLimit time.Duration = time.Second * 60 * 15 / 180

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
		panic(fmt.Sprintf("%s: %s", msg, err))
	}
}

func connectDatadog() *statsd.Client {
	c, err := statsd.New("datadog:8125")
	failOnError(err, "Failed to open connection to DogStatsD")
	c.Namespace = "dockercon2017golang."
	c.Tags = append(c.Tags, coolTagForApp)
  fmt.Println("datadog connected")
	return c
}

func collectTweets(subject string, newestId int64) anaconda.SearchResponse {
  fmt.Println("collecting tweets")
	api := anaconda.NewTwitterApi(twitterAccessToken, twitterTokenSecret)
	options := url.Values{}
	options.Set("count", "100")
	options.Set("result_type", "recent")
	options.Set("max_id", strconv.FormatInt(newestId, 10))
	searchResult, _ := api.GetSearch(subject, options)
	datadogClient := connectDatadog()
	datadogClient.Count("app.tweetscollected", 100, nil, apiRateLimit.Seconds())
	defer datadogClient.Close()
	return searchResult
}

func main() {
  fmt.Println("starting app")
  host, _ := os.Hostname()

	datadogClient := connectDatadog()
	defer datadogClient.Close()
	err := datadogClient.Event(&statsd.Event{
		fmt.Sprintf("TwitterGoLang Started on %v", host), //title
		"Started Twitter GoLang App for DockerCon 2017", //text
		time.Now(),     //time
		host,           //hostname
		aggregationKey, //aggregation key
		statsd.Normal,  //priority , Low or Normal
		aggregationKey, //sourcetypename
		statsd.Info,    //event type, Info, Error, Warning, or Success
		[]string{"lang:golang", "demo:dockercon"},
	})
	rabbitconn, err := amqp.Dial("amqp://guest:guest@rabbit:5672")
  failOnError(err, "Failed to connect to Rabbit")
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
	ch.QueueDeclare("TweetQ", false, false, false, false, nil)
	err = ch.QueueBind("TweetQ", "logstash", "tweets", false, nil)
	failOnError(err, "Failed to declare queue")
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Hello Web")
  })
  http.ListenAndServe(":8123", nil)
  anaconda.SetConsumerKey(twitterConsumerKey)
	anaconda.SetConsumerSecret(twitterConsumerSecret)
	maxId := int64(9223372036854775807)
	limiter := time.Tick(apiRateLimit)
	for {
		<-limiter
		tweets := collectTweets("dockercon OR docker OR datadog", maxId)

		for _, tweet := range tweets.Statuses {
			tweettimestamp, _ := time.Parse(time.RubyDate, tweet.CreatedAt)
			body := []byte(fmt.Sprintf("%v;;;%v;;;%v;;;%v", tweettimestamp.Format(time.RFC3339Nano), strconv.FormatInt(tweet.Id, 10), tweet.User.ScreenName, tweet.Text))
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
					Timestamp:   tweettimestamp,
				})
			failOnError(err, "Failed to publish a message")
		}
	}
}
