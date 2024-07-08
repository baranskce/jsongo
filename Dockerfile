FROM golang:1.24-alpine
WORKDIR /Users/baran/Desktop/bts/go-project

COPY go.mod go.sum ./
RUN go mod download

COPY *.go ./

RUN go build -o /Users/baran/Desktop/bts/go-project/main .

EXPOSE 8080

CMD ["./main"]