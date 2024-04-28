package driver

import (
	cf "github.com/backend-magang/cats-social-media/config"
	"github.com/jmoiron/sqlx"

	_ "github.com/lib/pq"

	"log"
)

func InitPostgres(config cf.Config) *sqlx.DB {
	psqlInfo := config.GetDSN()
	db, err := sqlx.Connect("postgres", psqlInfo)
	if err != nil {
		log.Fatal(err)
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
