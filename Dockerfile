FROM golang:1.21.1

RUN apt-get update && apt-get install -y build-essential make

RUN go install github.com/ramya-rao-a/go-outline@latest
RUN go install golang.org/x/tools/gopls@latest

WORKDIR /app
