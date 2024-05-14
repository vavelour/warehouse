package repository

import (
	"context"
	"fmt"
	"log"

	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/vavelour/warehouse/warehouse/internal/config"
)

const (
	maxConns = 10
)

type DB struct {
	DB *pgxpool.Pool
}

func NewDB(cfg config.DBConfig) (*DB, error) {
	poolConfig, err := pgxpool.ParseConfig(fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=%s",
		cfg.Host, cfg.Port, cfg.User, cfg.DBName, cfg.Password, cfg.SSLMode))
	if err != nil {
		return nil, fmt.Errorf("failed to parse database config: %w", err)
	}

	poolConfig.MaxConns = maxConns

	pool, err := pgxpool.ConnectConfig(context.Background(), poolConfig)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %w", err)
	}

	log.Println("Connect: OK.")

	if err := pool.Ping(context.Background()); err != nil {
		pool.Close()
		return nil, fmt.Errorf("failed to ping database: %w", err)
	}

	log.Println("Ping: OK.")

	return &DB{
		DB: pool,
	}, nil
}
