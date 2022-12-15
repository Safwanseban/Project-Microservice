package configs

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"time"
)

var count int64

func OpenDB(dsn string) (*sql.DB, error) {
	db, err := sql.Open("pgx", dsn)
	if err != nil {
		return nil, err
	}
	err = db.Ping()
	if err != nil {
		return nil, err
	}
	return db, nil
}
func ConnecTodb() *sql.DB {
	dsn := os.Getenv("DSN")
fmt.Println(dsn)
	for {
		connection, err := OpenDB(dsn)
		fmt.Println(connection)
		if err != nil {
			log.Println("cpostgres not yet connected")
			count++
		} else {
			log.Println("conncected to postgres")
		}
		if count > 10 {
			log.Println(err)
			return nil
		}
		log.Println("backiung of two seconds")
		time.Sleep(2 * time.Second)
		continue
	}
}
