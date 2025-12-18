package bootstrap

import (
	"GoTasks/task8/internal/rest/config"
	"database/sql"

	_ "github.com/jackc/pgx/stdlib"
)

func InitDB(cfg *config.Config) (*sql.DB, error) {
	db, err := sql.Open("pgx", cfg.PG)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}
