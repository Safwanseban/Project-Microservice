package main

import (
	"log"

	"github.com/Safwanseban/Project-Microservices/authentication-service/configs"
	"github.com/Safwanseban/Project-Microservices/authentication-service/routes"
	"github.com/gin-gonic/gin"
)

func init() {
	configs.GetEnv()
	configs.ConnecTodb()
}

var R = gin.Default()

type Config struct {
	R      *gin.Engine


}

func main() {
	log.Println("starting authentication service")
	routes.Routes(R)

	R.Run(":8086")
	// app := Config{
	// 	DB:     conn,
	// 	models: models.New(conn),
	// }
	// srv := &http.Server{
	// 	Addr:    ":8086",
	// 	Handler: app.R,
	// }

	// err := srv.ListenAndServe()
	// if err != nil {
	// 	log.Panic(err)
	// }

}
