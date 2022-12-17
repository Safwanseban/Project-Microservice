package configs

import (
	"log"

	"github.com/joho/godotenv"
)

func GetEnv() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("eror loading env")
	}
}
