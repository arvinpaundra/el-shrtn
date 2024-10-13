package database

import (
	"log"
	"sync"

	"github.com/arvinpaundra/el-shrtn/config"
	"go.mongodb.org/mongo-driver/mongo"
)

var (
	dbConn *mongo.Database
	dbErr  error
	once   sync.Once
)

func createConnection() {
	// Create database configuration information
	dbCfg := dbConfig{
		uri: config.GetMongoURI(),
	}

	// Create only one database Connection, not the same as database TCP connection
	once.Do(func() {
		dbConn, dbErr = dbCfg.connect()
		if dbErr != nil {
			log.Fatalf("failed connected to database: %s", dbErr.Error())
		}
	})

	log.Println("connected to database")
}

func GetConnection() *mongo.Database {
	// Check db connection, if exist return the memory address of the db connection
	if dbConn == nil {
		createConnection()
	}
	return dbConn
}
