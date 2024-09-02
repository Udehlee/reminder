package db

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/jackc/pgx/v5/stdlib"
	"github.com/pressly/goose/v3"
)

type PgConn struct {
	Conn *pgxpool.Pool
}

// connectDB establish a connection to the database
func connectDB(config Config) (*pgxpool.Pool, error) {
	ctx := context.Background()

	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		config.Host, config.Port, config.Username,
		config.Password, config.DbName)

	pool, err := pgxpool.New(ctx, dsn)
	if err != nil {
		return nil, fmt.Errorf("unable to create connection pool: %v", err)
	}

	db := stdlib.OpenDBFromPool(pool)

	if err := Migrations(db); err != nil {
		db.Close() // Close the connection if migrations fail
		return nil, fmt.Errorf("migration failed: %v", err)
	}

	return pool, nil
}

func Migrations(db *sql.DB) error {

	if err := goose.SetDialect("postgres"); err != nil {
		return fmt.Errorf("failed to set dialect: %v", err)
	}

	if err := goose.Up(db, "db/migrations"); err != nil {
		return fmt.Errorf("failed to apply migrations: %v", err)
	}

	return nil
}

func InitPG(cfg Config) (*PgConn, error) {
	pg := PgConn{}

	db, err := connectDB(cfg)
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	pg.Conn = db
	return &pg, nil
}
