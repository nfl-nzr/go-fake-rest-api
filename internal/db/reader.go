package db

import (
	"encoding/json"
	"errors"
	"io/fs"
	"io/ioutil"
	"os"

	"github.com/gabriel-vasile/mimetype"
)

func (d *Database) LoadFile(dsn string) error {
	var data, err = parseFile(dsn)
	if err != nil {
		return err
	}
	d.Data = data
	return nil
}

func FileValid(path string) (fs.FileInfo, error) {
	fi, err := os.Stat(path)
	if err != nil {
		return nil, err
	}

	mtype, err := mimetype.DetectFile(path)

	if err != nil {
		return nil, err
	}
	
	if mtype.String() != "application/json" {
		return nil, errors.New("file is either corrupt or it is not a json")
	}

	isDir := fi.IsDir()
	if isDir {
		return nil, errors.New("path is a directory")
	}

	return fi, nil
}

func FolderValid (path string) error {
	fi, err := os.Stat(path)
	if err != nil {
		return err
	}
	isDir := fi.IsDir()
	if isDir {
		return err
	}
	return nil 
}

func parseFile(dsn string) (*map[string]interface{}, error) {

	jsonFile, err := os.Open(dsn)

	if err != nil {
		return nil, err
	}

	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)

	var result map[string]interface{}
	
	err = json.Unmarshal([]byte(byteValue), &result)

	if err != nil {
		return nil, errors.New("file is corrupt/invalid")
	}

	return &result, nil
}
