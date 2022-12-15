package main

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/Safwanseban/Project-Microservices/authentication-service/configs"
	"github.com/Safwanseban/Project-Microservices/authentication-service/models"
	"github.com/gin-gonic/gin"
	_ "github.com/jackc/pgconn"
	_ "github.com/jackc/pgx/stdlib"
	_ "github.com/jackc/pgx/v4"
)

var R = gin.Default()

type Config struct {
	R *gin.Engine
	DB     *sql.DB
	models models.Models
}

func main() {
	log.Println("starting authentication service")

	conn := configs.ConnecTodb()
	if conn == nil {
		log.Panic("error connceting to db")
	}
	app := Config{
		DB:     conn,
		models: models.New(conn),
	}
	srv := &http.Server{
		Addr:    ":8086",
		Handler: app.R,
	}

	err:=srv.ListenAndServe()
	if err != nil {
		log.Panic(err)
	}
}
