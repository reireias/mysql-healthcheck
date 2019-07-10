FROM golang:alpine

WORKDIR /go/src/app
COPY . .
COPY ./docker-entrypoint.sh /
RUN apk add git
RUN go get -d -v ./...
RUN go build -o main main.go

ENTRYPOINT [ "/docker-entrypoint.sh" ]
