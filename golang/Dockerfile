FROM golang:1.7-alpine
EXPOSE 8080 8123

RUN apk --update add git wget
COPY . /go/src/app
WORKDIR /go/src/app
RUN go get -d
RUN go install -v

RUN wget https://github.com/jwilder/dockerize/releases/download/v0.1.0/dockerize-linux-amd64-v0.1.0.tar.gz
RUN tar -C /usr/local/bin -xzvf dockerize-linux-amd64-v0.1.0.tar.gz

CMD ["dockerize", "-wait", "http://rabbit:15672", "-timeout", "20s", "app"]
