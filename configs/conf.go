package configs

import (
	"log"

	"github.com/joho/godotenv"
)

func RunProject() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	ConnectDataBase()
}

func RunTestProject(envFile string) {
	err := godotenv.Load(envFile)
	if err != nil {
		log.Fatal("Error loading .env.test file")
	}
	ConnectDataBase()
}
