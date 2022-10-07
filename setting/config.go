package setting

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	SOCKET_BINANCE_URL string
}

var C *Config

func LoadConfig() {

	err := godotenv.Load("./setting/.env")
	if err != nil {
		log.Fatal("Error load .env file")
	}

	C = &Config{
		SOCKET_BINANCE_URL: os.Getenv("SOCKET_BINANCE_URL"),
	}
}
