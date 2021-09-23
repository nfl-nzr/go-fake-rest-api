package db

import (
	"encoding/json"
	"io/ioutil"
	"log"
)

func (d *Database) WriteToDB(filePath string) error {
	b, err := json.MarshalIndent(d.Data, "", "    ")
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(filePath, b, 0666)
	if err != nil {
		return err
	}
	log.Println("Wrote db to file: ", filePath)
	return nil
}
