FROM golang:1.21.6-alpine3.19 AS builder

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /app/server ./cmd/api

FROM alpine:3.19

WORKDIR /app
COPY --from=builder /app/server /app/server

EXPOSE 8080
ENV PORT=8080

CMD ["/app/server"]
