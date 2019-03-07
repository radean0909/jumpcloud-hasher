package database

import (
	"errors"
	"io/ioutil"
	"log"
	"path/filepath"
	"strings"
)

type DB struct {
	Name   string
	Tables map[string]*DBTable
}

type DBTable struct {
	Data map[int64]string
	Meta *DBMeta
}

type DBMeta struct {
	Increment int64
	Count     int64
	Hashtime  float32
}

func NewDBMeta() *DBMeta {
	return &DBMeta{
		Increment: 1,
		Count:     0,
		Hashtime:  0.0,
	}
}

func Connect(dbName string) (*DB, error) {
	/*
	 * Create a new "database", in the real world this would connect to a database using the database driver and database/sql library
	 * Since external libraries aren't allowed, we are faking it here with an in-memory map, the key is a string representing the table-name with the value being a DB object
	 */

	if dbName == "" {
		return nil, errors.New("No DB specified")
	}

	db := &DB{
		Name:   dbName,
		Tables: make(map[string]*DBTable),
	}

	// Create a table for each of our models
	files, err := ioutil.ReadDir("./models")

	if err != nil {
		log.Printf("Error opening models, %v", err.Error())
		return nil, err
	}

	for _, file := range files {
		// If this is a proper model file, create the "table" in the database, in the real world, these models would contain schema for the tables, likely
		if filepath.Ext(file.Name()) == ".model" {
			tableName := strings.Replace(filepath.Base(file.Name()), ".model", "", 1)
			db.Tables[tableName] = &DBTable{
				Data: make(map[int64]string),
				Meta: NewDBMeta(),
			}
		}
	}

	return db, nil

}

func (table *DBTable) Update(id int64, val string) {
	table.Data[id] = val
}
