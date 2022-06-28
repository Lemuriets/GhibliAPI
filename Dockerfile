FROM golang:1.17

RUN apt update
RUN apt install -y git

RUN mkdir -p /app
WORKDIR /app

COPY . .

RUN go get -u
