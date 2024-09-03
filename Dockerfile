FROM golang:1.23 AS builder

WORKDIR /app

COPY go.mod ./
RUN go mod download

COPY . .

RUN go build -o /go-api ./cmd

FROM alpine:latest

WORKDIR /root/

COPY --from=builder /go-api .

ENTRYPOINT ["./go-api"]

EXPOSE 8080
