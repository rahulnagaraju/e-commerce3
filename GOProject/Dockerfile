FROM golang:alpine

RUN mkdir /ecommerce1

WORKDIR /ecommerce1

ADD go.mod .
ADD go.sum .

RUN go mod download
ADD . .

RUN go get github.com/githubnemo/CompileDaemon

EXPOSE 9012


ENTRYPOINT CompileDaemon --build="go build main.go" --command=./main
