package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

var ADDR = "ADDR"

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	cfg := config{
		addr: os.Getenv(ADDR),
	}

	app := &application{
		cfg,
	}

	mux := app.mount()
	log.Fatal(app.run(mux))
}
