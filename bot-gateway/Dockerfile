FROM golang:alpine as modules
COPY go.mod go.sum /go/trend-bot/bot-gateway/modules/
WORKDIR /go/trend-bot/bot-gateway/modules
RUN go mod download

FROM golang:alpine as builder
COPY --from=modules /go/pkg /go/pkg
WORKDIR /github.com/IamVladlen/trend-bot/bot-gateway
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o .bin/bot/ ./cmd/bot/main.go

FROM alpine
WORKDIR /trend-bot/bot-gateway
COPY --from=builder /github.com/IamVladlen/trend-bot/bot-gateway/.bin/bot .
COPY --from=builder /github.com/IamVladlen/trend-bot/bot-gateway/config config/
CMD ["./main"]