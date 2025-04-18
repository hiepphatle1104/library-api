package database

import (
	"context"
	"database/sql"
	"library-api/model"

	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
	"github.com/uptrace/bun/driver/pgdriver"
)

type Database struct {
	DB *bun.DB
}

func NewConnection() (*Database, error) {
	dsn := "postgres://postgres:admin@localhost:5432/librarydb?sslmode=disable"
	db := bun.NewDB(sql.OpenDB(pgdriver.NewConnector(pgdriver.WithDSN(dsn))), pgdialect.New())

	err := db.Ping()
	if err != nil {
		return nil, err
	}

	// db.AddQueryHook(bundebug.NewQueryHook(
	// 	bundebug.WithVerbose(true),
	// 	bundebug.FromEnv("BUNDEBUG"),
	// ))

	ctx := context.Background()
	newTable[model.User](db, ctx)

	return &Database{DB: db}, nil
}

func (db *Database) Close() error {
	return db.DB.Close()
}

func newTable[T any](db *bun.DB, ctx context.Context) {
	_, err := db.NewCreateTable().Model((*T)(nil)).IfNotExists().Exec(ctx)
	if err != nil {
		panic(err)
	}
}
