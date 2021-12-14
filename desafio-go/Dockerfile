FROM golang:1.17

WORKDIR /go/src
COPY . .
RUN apt-get update && apt-get install sqlite3 -y
RUN go build -o /go/bin/gateway cmd/main.go 

CMD ["tail", "-f", "/dev/null"]