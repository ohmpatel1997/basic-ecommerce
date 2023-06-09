FROM  golang:1.19-alpine3.15 AS builder

ENV GO111MODULE=on

WORKDIR /code


COPY go.mod .
COPY go.sum .
RUN go mod download

COPY . .

# build the binary
RUN env CGO_ENABLED=0 GOOS=linux  go build -o /api cmd/main.go

# final stage
FROM alpine:3.17.3

RUN apk add curl
COPY --from=builder /api /
COPY pkg/migrations /migrations/
COPY wait-for.sh /


RUN chmod +x /api
