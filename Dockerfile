FROM golang:1.24-alpine AS builder

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -o flylang .

FROM alpine:latest

ENV SERVER_NAME=americandes
ENV APP_PORT=19001

WORKDIR /app
COPY --from=builder /app/flylang .

EXPOSE ${APP_PORT}

ENTRYPOINT ["sh", "-c", "./flylang --server=${SERVER_NAME}"]
