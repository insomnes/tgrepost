FROM golang:1.23.4-alpine AS builder

RUN apk add --no-cache git
WORKDIR /tgreposter

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o tgreposter .

# Runner
FROM alpine:3.21
RUN apk add --no-cache ca-certificates
WORKDIR /tgreposter
COPY --from=builder /tgreposter/tgreposter .

CMD ["./tgreposter"]
