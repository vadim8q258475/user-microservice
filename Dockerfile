FROM golang:1.23-alpine as builder

WORKDIR /app

COPY . .

RUN go mod tidy

RUN go build -o main

FROM alpine:latest

COPY --from=builder /app/main /main

EXPOSE 1234

CMD ["/main"]