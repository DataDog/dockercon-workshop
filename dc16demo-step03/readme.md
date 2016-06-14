Add a golang app to push tweets into RabbitMQ

Login to Twitter.com, navigate to https://apps.twitter.com/ and create an app.  Find the link to *manage keys and access tokens*. Click the Create my Tokens button.

Create environment variables in your shell and assign the appropriate values for:

    twitterConsumerKey
    twitterConsumerSecret
    twitterAccessToken
    twitterTokenSecret

Review golang/main.go

Review golang/Dockerfile

    docker-compose up

Notice the potential problem. Start order. Health check is coming soon in Docker, but until then look at Dockerize (https://github.com/jwilder/dockerize).

Review golang/Dockerfile2

Review docker-compose2.yaml that refers to the new Dockerfile2 file

    docker-compose -f docker-compose2.yaml up

Go to http://localhost:8080 and look around the rabbit interface. Review the Exchange, queue defined in the go app.

