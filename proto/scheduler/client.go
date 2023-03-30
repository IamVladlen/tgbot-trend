package grpcscheduler

import (
	"log"

	grpc "google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type Client struct {
	SchedulerClient
}

func New(uri string) *Client {
	conn, err := grpc.Dial(uri, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal(err)
	}

	client := NewSchedulerClient(conn)

	return &Client{
		client,
	}
}
