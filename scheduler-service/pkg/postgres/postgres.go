package postgres

import (
	"context"
	"log"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

type DB struct {
	*pgxpool.Pool
}

// New creates new Postgres connection pool.
func New(url string) *DB {
	pg := &DB{}

	// Initialize config for postgres server
	pgconfig, err := pgxpool.ParseConfig(url)
	if err != nil {
		log.Fatalln("Can't connect to postgres:", err)
	}

	// Try to create connection pool
	for i := 10; i > 0; i-- {
		pg.Pool, err = pgxpool.NewWithConfig(context.Background(), pgconfig)
		if err == nil {
			break
		}

		log.Printf("trying to connect to postgres, attempts left: %d\n", i)

		<-time.After(1 * time.Second)
	}

	// Health check
	if err := pg.Ping(context.Background()); err != nil {
		log.Fatalln("Cannot connect to postgres:", err)
	}

	return pg
}

// Close gracefully closes all connections in the pool.
func (p *DB) Close() {
	p.Pool.Close()
}
