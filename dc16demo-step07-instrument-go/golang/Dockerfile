FROM golang:1.7-alpine
EXPOSE 8080 8123

RUN apk --update add git wget
COPY . /go/src/app
WORKDIR /go/src/app
RUN go get -d
RUN go install -v

ENV DOCKERIZE_VERSION v0.2.0
RUN wget https://github.com/jwilder/dockerize/releases/download/$DOCKERIZE_VERSION/dockerize-linux-amd64-$DOCKERIZE_VERSION.tar.gz \
    && tar -C /usr/local/bin -xzvf dockerize-linux-amd64-$DOCKERIZE_VERSION.tar.gz

CMD ["dockerize", "-wait", "http://rabbit:15672", "-timeout", "20s", "app"]
