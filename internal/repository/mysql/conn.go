package mysql

import (
	"context"
	"database/sql"
	"go-clean-example/config"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

func Open() (*sql.DB, error) {
	dsn := config.GetConfig().DSN
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := db.PingContext(ctx); err != nil {
		_ = db.Close()
		return nil, err
	}
	db.SetMaxOpenConns(25)
	db.SetMaxIdleConns(10)
	db.SetConnMaxLifetime(5 * time.Minute)
	db.SetConnMaxIdleTime(5 * time.Minute)
	return db, nil
}
