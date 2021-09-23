package main

import (
	"errors"
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/nfl-nzr/go-fake-rest-api/internal/db"
	"github.com/nfl-nzr/go-fake-rest-api/internal/server"
)

//Exit code in case server goes belly up.
const exitFail = 1
const version =  "0.0.1"

var Usage = func() {
        fmt.Fprintf(os.Stdout, "Usage of fake-rest-server v%s\n", version)
        flag.PrintDefaults()
}

func run() error {
	var config = server.Config{}
	flag.Usage = Usage
	flag.IntVar(&config.Port, "port", 4000, "Port to run the server on")
	flag.StringVar(&config.FilePath, "file", "", "Path to db file")
	flag.StringVar(&config.Env, "env", "prod", "Environment. Possible values { prod | dev }")
	flag.BoolVar(&config.ReadOnlyMode, "r", true, "Prevent writes to the file. Possible values {true | false} ")
	flag.StringVar(&config.StaticFiles, "serve-static", "", "Serve satic files from dir specified (optional")
	flag.Parse()

	// Exit if db file is empty
	if config.FilePath == "" {
		fmt.Println("file cannot be empty. exiting...")
		os.Exit(exitFail)
	}

	_, err := db.FileValid(config.FilePath)
	if err != nil {
		return err
	}

	err = db.FolderValid(config.StaticFiles)

	if err != nil && config.StaticFiles != "" {
		if os.IsNotExist(err) {
			err = os.MkdirAll(config.StaticFiles, os.ModePerm)
			if err != nil {
				return errors.New("error making the directory")
			}
		} else  {
			return errors.New("assets path is not valid")
		}
	}

	var app = server.Application{Cfg: config}
	app.CreateServer()
	err = app.StartServer()
	if err == nil {
		return nil
	} else {
		return err
	}
}

func main() {
	if err := run(); err != nil {
		log.Printf("Error: %s\n", err)
		os.Exit(exitFail)
	}
}
