FROM golang:tip-alpine3.22 AS builder

WORKDIR /app

COPY back/go.mod ./
RUN go mod tidy

COPY back/ ./back
COPY ihm/ ./ihm

WORKDIR /app/back
RUN go build -o /usr/local/bin/server main.go

FROM alpine:latest

WORKDIR /app

COPY --from=builder /usr/local/bin/server .
COPY --from=builder /app/ihm ./ihm

EXPOSE 8080

CMD ["./server"]