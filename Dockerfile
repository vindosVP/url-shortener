FROM golang:1.20.0

WORKDIR /usr/src/app
COPY . .
RUN go mod tidy
RUN go build -o url-shortener src/cmd/app/main.go
ENTRYPOINT ["./url-shortener"]