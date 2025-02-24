package database

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/doug-martin/goqu/v9"
	"github.com/lcv-back/goload/internal/configs"
)

func InitializeDB(databaseConfig configs.Database) (db *sql.DB, cleanup func(), err error) {
	connectionString := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s",
		databaseConfig.Username,
		databaseConfig.Password,
		databaseConfig.Host,
		databaseConfig.Port,
		databaseConfig.Database,
	)

	db, err = sql.Open("mysql", connectionString)
	if err != nil {
		log.Printf("error connecting to the database: %w\n", err)
		return nil, nil, err
	}

	cleanup = func() {
		db.Close()
	}

	return db, cleanup, nil
}

func InitializeGoquDB(db *sql.DB) *goqu.Database {
	return goqu.New("mysql", db)
}
