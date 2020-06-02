package database

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Config struct {
	URL  string
	Name string
}

const defaultURL = "mongodb://root:root@localhost:27017"
const defaultName = "codeinquarentena"

type DatabaseImpl struct {
	conn *mongo.Database
}

func New(config Config) *DatabaseImpl {
	url := config.URL
	name := config.Name

	if url == "" {
		url = defaultURL
	}
	if name == "" {
		name = defaultName
	}

	opts := options.Client().ApplyURI(url)
	conn, err := mongo.NewClient(opts)
	if err != nil {
		log.Fatalln("error on connection", err)
	}

	if err := conn.Connect(context.Background()); err != nil {
		log.Fatalln("error on connection", err)
	}

	return &DatabaseImpl{
		conn: conn.Database(name),
	}
}

func (d DatabaseImpl) Create(collection string, document interface{}) (interface{}, error) {
	res, err := d.conn.Collection(collection).InsertOne(nil, document)
	if err != nil {
		log.Fatalln("error on create", err)
		return nil, err
	}
	return res.InsertedID, nil
}

func (d DatabaseImpl) CreateMany(collection string, documents []interface{}) (interface{}, error) {
	res, err := d.conn.Collection(collection).InsertMany(nil, documents)
	if err != nil {
		log.Fatalln("error on insertmany", err)
		return nil, err
	}
	return res.InsertedIDs, nil
}
