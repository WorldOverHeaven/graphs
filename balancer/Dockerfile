FROM golang:1.20-alpine AS builder

WORKDIR /build

COPY . .

RUN go mod download

RUN go build -o main balancer/cmd/balancer/main.go

FROM alpine

WORKDIR /build

COPY --from=builder /build/main /build/main

COPY /deploy /build/deploy

CMD ["./main", "--rootPath", "."]