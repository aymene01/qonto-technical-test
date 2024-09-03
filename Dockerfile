FROM golang:1.23 AS builder

WORKDIR /app

COPY go.mod ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /go-api cmd/main.go cmd/routes.go

FROM alpine:latest

WORKDIR /root/

COPY --from=builder /go-api .

ENTRYPOINT ["./go-api"]

EXPOSE 8080
