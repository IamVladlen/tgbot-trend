FROM golang:alpine as modules
COPY go.mod go.sum /go/trend-bot/scheduler-service/modules/
WORKDIR /go/trend-bot/scheduler-service/modules
RUN go mod download

FROM golang:alpine as builder
COPY --from=modules /go/pkg /go/pkg
WORKDIR /github.com/IamVladlen/trend-bot/scheduler-service
COPY . .
RUN go build -o .bin/scheduler/ ./cmd/scheduler/main.go

FROM alpine
WORKDIR /trend-bot/scheduler-service
COPY --from=builder /github.com/IamVladlen/trend-bot/scheduler-service/.bin/scheduler .
COPY --from=builder /github.com/IamVladlen/trend-bot/scheduler-service/config config/
COPY --from=builder /github.com/IamVladlen/trend-bot/scheduler-service/migrations migrations/
CMD ["./main"]