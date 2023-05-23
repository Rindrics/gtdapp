# build
FROM golang:1.19 as builder

RUN mkdir /app
WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o /app/main ./cmd

# run
FROM alpine:latest

RUN apk --no-cache add ca-certificates

WORKDIR /root/

COPY --from=builder /app/main .

EXPOSE 8888

CMD ["./main"]