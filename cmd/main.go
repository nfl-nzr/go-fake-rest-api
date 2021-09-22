package main

import (
	"flag"
	"log"
	"os"

	"github.com/nfl-nzr/go-fake-rest-api/internal/db"
	"github.com/nfl-nzr/go-fake-rest-api/internal/server"
)

//Exit code in case server goes belly up.
const exitFail = 1

func run() error {
	var config = server.Config{}
	flag.IntVar(&config.Port, "port", 4000, "Port to run the server on")
	flag.StringVar(&config.FilePath, "file", "", "Path to db file")
	flag.StringVar(&config.Env, "env", "prod", "Environment { prod | dev }")
	flag.BoolVar(&config.ReadOnlyMode, "r", true, "Prevent writes to the file")

	// Exit if db file is empty
	if config.FilePath == "" {
		panic("ERRRRRRR")
	}

	flag.Parse()
	_, err := db.FileValid(config.FilePath)
	if err != nil {
		return err
	}

	

	var app = server.Application{Cfg: config}
	app.CreateServer()
	return app.StartServer()
}

func main() {
	if err := run(); err != nil {
		log.Printf("Error: %s\n", err)
		os.Exit(exitFail)
	}
}
