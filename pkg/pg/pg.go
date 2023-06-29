package pg_conn

import (
	"BalanceRange/internal/config"
	"fmt"

	"github.com/jmoiron/sqlx"
)

func NewPgClient(config config.PgConfig) (database *sqlx.DB, err error) {
	connectionURL := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s",
		config.Host, config.Port, config.User, config.Pwd, config.Db,
	)

	database, err = sqlx.Open("pgx", connectionURL)
	if err != nil {
		return nil, err
	}

	if err = database.Ping(); err != nil {
		return nil, err
	}

	return database, nil
}
