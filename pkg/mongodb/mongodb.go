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

type Deps struct {
	URI      string
	Username string
	Password string
	DBName   string
}

func New(d Deps) *DB {
	ctx, cancel := context.WithTimeout(context.Background(), _connectionTimeout*time.Second)
	defer cancel()

	opts := options.Client().ApplyURI(d.URI)
	opts.SetAuth(options.Credential{
		Username: d.Username,
		Password: d.Password,
	})

	clt, err := mongo.Connect(ctx, opts)
	if err != nil {
		log.Fatalln("Can't connect to MongoDB:", err)
	}

	if err := clt.Ping(context.Background(), nil); err != nil {
		log.Fatalln("Can't connect to MongoDB:", err)
	}

	db := clt.Database(d.DBName)

	return &DB{
		db,
	}
}

func (db *DB) Disconnect() error {
	ctx, cancel := context.WithTimeout(context.Background(), _connectionTimeout)
	defer cancel()
	if err := db.Client().Disconnect(ctx); err != nil {
		return err
	}

	return nil
}