FROM golang:1.17

WORKDIR /go/src

RUN apt-get update && apt-get install sqlite3 -y

CMD ["tail", "-f", "/dev/null"]