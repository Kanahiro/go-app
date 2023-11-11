package infra

import (
	"fmt"
	"log"
	"os"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

var schema = `
CREATE TABLE IF NOT EXISTS todo (
		id serial PRIMARY KEY,
		text text NOT NULL,
		done boolean NOT NULL
);`

var DbInstance *sqlx.DB

func InitDb() *sqlx.DB {
	if DbInstance != nil {
		return DbInstance
	}

	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Tokyo",
		os.Getenv("POSTGRES_HOST"),
		os.Getenv("POSTGRES_USER"),
		os.Getenv("POSTGRES_PASSWORD"),
		os.Getenv("POSTGRES_DB"),
		os.Getenv("POSTGRES_PORT"),
	)
	DbInstance, err := sqlx.Connect("postgres", dsn)

	DbInstance.MustExec(schema)

	if err != nil {
		log.Fatal("failed to init database: ", err)
	}

	return DbInstance
}

func Db() *sqlx.DB {
	return InitDb()
}
