FROM golang:alpine as builder

WORKDIR /github.com/IamVladlen/trend-bot/bot-gateway
COPY . .
RUN go build -o .bin/bot/ ./cmd/bot/main.go

FROM alpine

WORKDIR /trend-bot/bot-gateway
COPY --from=builder /github.com/IamVladlen/trend-bot/bot-gateway/.bin/bot .
COPY --from=builder /github.com/IamVladlen/trend-bot/bot-gateway/config config/
CMD ["./main"]