package topic

import (
	"fmt"
	"log"

	"github.com/joho/godotenv"
)

func init() {
	var err error
	//load .env file
	err = godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	fmt.Println(exchangeName)
	fmt.Println(typeName)
}
