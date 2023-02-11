package mongodb

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	_connectionTimeout = 5
)

type DB struct {
	*mongo.Database
}

func New(uri string, username, password string, dbName string) *DB {
	ctx, cancel := context.WithTimeout(context.Background(), _connectionTimeout*time.Second)
	defer cancel()

	opts := options.Client().ApplyURI(uri)
	opts.SetAuth(options.Credential{
		Username: username,
		Password: password,
	})

	clt, err := mongo.Connect(ctx, opts)
	if err != nil {
		log.Fatalln("Can't connect to MongoDB:", err)
	}

	if err := clt.Ping(context.Background(), nil); err != nil {
		log.Fatalln("Can't connect to MongoDB:", err)
	}

	db := clt.Database(dbName)

	return &DB{
		db,
	}
}
