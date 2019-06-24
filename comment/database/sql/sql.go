package sql

import (
	"fmt"
	"sync"

	"github.com/jinzhu/gorm"

	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var (
	client *gorm.DB
	once   sync.Once
)

func Init(dbHost string, dbPort int, dbUser, dbPass, dbName string, debug bool) *gorm.DB {
	var err error
	once.Do(func() {
		dbConfig := fmt.Sprintf("host=%s port=%d user=%s dbname=%s password=%s sslmode=disable", dbHost, dbPort, dbUser, dbName, dbPass)
		client, err = gorm.Open("postgres", dbConfig)
		if err != nil {
			fmt.Println(err)
		}
		if debug {
			client.LogMode(true)
		}
		client.SingularTable(true)
	})
	return client
}

func GetInstance() *gorm.DB {
	if client == nil {
		fmt.Println("DB is not initialized. Please execute Init first")
		return nil
	}
	return client
}
