package database

import (
	"context"

	"github.com/arvinpaundra/el-shrtn/config"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Define the database conn configuration
type (
	dbConfig struct {
		uri string
	}
)

// Connect to postgres with the input configuration
func (conf dbConfig) connect() (*mongo.Database, error) {
	opt := options.Client().ApplyURI(conf.uri)

	conn, err := mongo.Connect(context.TODO(), opt)
	if err != nil {
		return nil, err
	}

	db := conn.Database(config.GetMongoDBName())

	return db, nil
}
