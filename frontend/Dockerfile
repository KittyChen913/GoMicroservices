FROM golang:1.24-alpine AS builder

WORKDIR /app

COPY go.mod ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o ./frontend .

FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/frontend .
COPY --from=builder /app/index.gohtml .

CMD ["./frontend"]