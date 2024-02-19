package db

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5"
	"os"
	"time"
)

type DatabaseAdapter interface {
	Connect() error
	AddKey(secret string) error
}

type Key struct {
	Id       int
	Secret   string
	IssuedAt int64
}

type TimeRetriever func() int64

type DatabasePostgres struct {
	connection    string
	timeRetriever TimeRetriever
}

func NewDatabasePostgres() *DatabasePostgres {
	return &DatabasePostgres{
		connection:    "postgres://postgres:postgres@localhost:5432/sciffr?sslmode=disable",
		timeRetriever: now,
	}
}

func now() int64 {
	return time.Now().Unix()
}

func (d *DatabasePostgres) Connect() error {
	conn, err := pgx.Connect(context.Background(), d.connection)
	if err != nil {
		return err
	}
	defer func(conn *pgx.Conn, ctx context.Context) {
		_ = conn.Close(ctx)
	}(conn, context.Background())
	err = initialize(conn)
	if err != nil {
		return err
	}
	return nil
}

func initialize(conn *pgx.Conn) error {
	dir, err := os.ReadDir("sql")
	if err != nil {
		return err
	}
	var sqls []string
	for _, file := range dir {
		var readFile []byte
		readFile, err = os.ReadFile("sql/" + file.Name())
		if err != nil {
			return err
		}
		sqls = append(sqls, string(readFile))
	}
	tx, err := conn.BeginTx(context.Background(), pgx.TxOptions{})
	if err != nil {
		return err
	}
	for _, sql := range sqls {
		fmt.Println("Executing: ", sql)
		_, err = tx.Exec(context.Background(), sql)
		if err != nil {
			_ = tx.Rollback(context.Background())
			return err
		}
	}
	_ = tx.Commit(context.Background())
	return nil
}

func (d *DatabasePostgres) AddKey(secret string) error {
	conn, err := pgx.Connect(context.Background(), d.connection)
	if err != nil {
		return err
	}
	defer func(conn *pgx.Conn, ctx context.Context) {
		_ = conn.Close(ctx)
	}(conn, context.Background())
	_, err = conn.Exec(context.Background(),
		"INSERT INTO keys (secret, issued_at) VALUES ($1, $2)", secret, d.timeRetriever())
	if err != nil {
		return err
	}
	return nil
}
