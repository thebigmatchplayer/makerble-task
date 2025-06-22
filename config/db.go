package config

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/lib/pq"
	"go.uber.org/zap"

	"github.com/joho/godotenv"
)

var DB *sql.DB

func InitDB() {
	_ = godotenv.Load()

	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s sslmode=require",
		os.Getenv("PGHOST"),
		os.Getenv("PGUSER"),
		os.Getenv("PGPASSWORD"),
		os.Getenv("PGDATABASE"),
	)

	var err error
	DB, err = sql.Open("postgres", dsn)
	if err != nil {
		Log.Fatal("Failed to open DB", zap.Error(err))
	}

	if err = DB.Ping(); err != nil {
		Log.Fatal("Failed to ping DB", zap.Error(err))
	}

	Log.Info("Connected to Neon PostgreSQL")
}
