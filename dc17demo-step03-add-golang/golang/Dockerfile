FROM golang:1.7-alpine

RUN apk --update add git curl

COPY . /go/src/app
WORKDIR /go/src/app
RUN go get -d
RUN go install -v

CMD ["app"]
