FROM golang:1.19-alpine as builder

COPY . /app
WORKDIR /app

RUN go mod download
RUN go build -o ./build/apiserv cmd/api/main.go

FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/build/apiserv .
COPY --from=builder /app/configs configs

EXPOSE 8000

CMD ["/app/apiserv"]