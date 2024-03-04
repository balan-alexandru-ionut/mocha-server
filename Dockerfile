FROM golang:alpine

RUN mkdir -p /mocha-config

WORKDIR /mocha

COPY . .

RUN go mod download
RUN go build main.go

EXPOSE 8080

CMD ["./main"]