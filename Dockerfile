FROM golang:1.13-alpine

RUN mkdir -p /app
WORKDIR /app

ADD . /app

ENV GO111MODULE=on
#development only
RUN ["go", "get", "github.com/githubnemo/CompileDaemon"]

#production
RUN GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" -o /dist/api /app/cmd/api/main.go
RUN GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" -o /dist/cron /app/cmd/cron/main.go

#development only
RUN apk add --update gcc make build-base

EXPOSE 8000

CMD ["/dist/api"]
