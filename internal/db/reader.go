package db

import (
	"errors"
	"io/fs"
	"os"

	"github.com/gabriel-vasile/mimetype"
)

func (d *Database) LoadFile(dsn string) (error) {
	
	var data = parseFile(dsn)
	d.data = data
	return nil
}

func FileValid (path string) (fs.FileInfo, error) {
	fi ,err := os.Stat(path);
	if err != nil {
		return  nil, err
	}
	
	mtype, err := mimetype.DetectFile(path)
	
	if err != nil {
		return nil, err
	}
	
	if mtype.String() != "application/json" {
		return nil, errors.New("not a json file")
	}
	
	isDir := fi.IsDir()
	if isDir {
		return nil, errors.New("path is a directory")
	}

 	return fi, nil
}

func parseFile(dsn string) *map[string]interface{} {
	var data = map[string]interface{} {
		"Key": "Value",
	}
	return &data
}

