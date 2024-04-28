package driver

import (
	"database/sql"

	_ "github.com/lib/pq"

	"log"

	"github.com/backendmagang/project-1/config"
)

func InitPostgres(cfg config.DBConfig) *sql.DB {
	psqlInfo := cfg.GetDSN()
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		log.Println(err)
		return nil
	}
	log.Println("[Database] initialized...")

	err = db.Ping()
	if err != nil {
		log.Println("[Database] failed to connect to database: ", err)
		return nil
	}

	log.Println("[Database] successfully connected")
	return db
}
