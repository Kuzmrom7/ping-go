FROM golang:latest

WORKDIR $GOPATH/app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -o build ./cmd

EXPOSE 8080

CMD ["./build"]