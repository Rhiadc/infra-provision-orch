FROM golang:alpine as builder

WORKDIR /src

COPY . .

RUN go mod download

RUN GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -o server

FROM alpine:latest

EXPOSE 8080

WORKDIR /app

COPY --from=builder /src/server .

ENTRYPOINT ["/app/server"]