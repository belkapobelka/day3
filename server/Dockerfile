# 1
FROM golang:alpine as builder

RUN apk update
RUN apk upgrade
RUN apk add --no-cache git

RUN mkdir /app

WORKDIR /app

COPY . .

COPY go.mod .
COPY go.sum .

RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o server

# 2
FROM alpine:latest

RUN apk --no-cache add ca-certificates

RUN mkdir /app

WORKDIR /app

COPY --from=builder /app/server .

CMD ["./server"]