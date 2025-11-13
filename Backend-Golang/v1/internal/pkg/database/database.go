package database

import (
	"fmt"

	"github.com/PPunyapatt/InventoryService-Vue-Golang/v1/internal/pkg/config"
	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/jmoiron/sqlx"
)

type Database struct {
	Sqlx *sqlx.DB
}

func InitDatabase(cfg *config.AppConfig) (*Database, error) {
	// -------------------- postgres --------------------
	dsn := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable", cfg.UserDB, cfg.PasswordDB, cfg.HostDB, cfg.PortDB, cfg.DatabaseName)
	dbSqlx, err := sqlx.Open("pgx", dsn)
	if err != nil {
		return nil, err
	}
	db := &Database{
		Sqlx: dbSqlx,
	}

	return db, nil
}
