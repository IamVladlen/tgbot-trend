FROM golang:alpine as builder

WORKDIR /github.com/IamVladlen/trend-bot
COPY . .
RUN go build -o .bin/bot/ ./cmd/bot/main.go

FROM alpine

WORKDIR /trend-bot/
COPY --from=builder /github.com/IamVladlen/trend-bot/.bin/bot .
COPY --from=builder /github.com/IamVladlen/trend-bot/config config/
COPY --from=builder /github.com/IamVladlen/trend-bot/.env .
CMD ["./main"]