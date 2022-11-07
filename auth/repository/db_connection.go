package repository

import (
	"log"
	"sync"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var once sync.Once

type connection gorm.DB

var (
	instance connection
)

func GetDBConnection() connection {
	once.Do(func() {
		conn, err := getConnection()
		if err != nil {
			log.Fatal("Database connection error")
		}
		instance = connection(*conn)
	})

	return instance
}

func getConnection() (*gorm.DB, error) {
	dsn := "host=localhost user=postgres password=postgres dbname=auth port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		return nil, err
	}
	return db, nil
}
