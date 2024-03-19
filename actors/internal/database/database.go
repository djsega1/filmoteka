package database

import (
	"context"
	"fmt"
	"log"

	"github.com/djsega1/filmoteka/actors/internal/config"
	"github.com/jackc/pgx/v5/pgxpool"
)

func NewPool(ctx context.Context, sc config.StorageConfig) (pool *pgxpool.Pool) {
	dsn := fmt.Sprintf("postgresql://%s:%s@%s:%s/%s", sc.Username, sc.Password, sc.Host, sc.Port, sc.Database)

	pool, err := pgxpool.New(ctx, dsn)
	if err != nil {
		log.Fatalf("Unable to create connection pool: %v\n", err)
	}

	log.Printf("Successfully connected to %s at %s:%s", sc.Database, sc.Host, sc.Port)

	return pool
}
