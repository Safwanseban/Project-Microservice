package initializers

import (
	"fmt"

	"github.com/joho/godotenv"
)

func GetEnv() {
	if err := godotenv.Load(); err != nil {
		fmt.Println("error loading env file")

	}

}
