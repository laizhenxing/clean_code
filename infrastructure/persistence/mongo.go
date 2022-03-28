package persistence

import (
	"context"
	"sync"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	MongoClient *mongo.Client
	mongoClient *mongo.Client
	once sync.Once
)

func InitEngine(uri string)  {
	var err error
	cliOptions := options.Client().ApplyURI(uri)

	// 连接到 mongo
	mongoClient, err = mongo.Connect(context.TODO(), cliOptions)
	if err != nil {
		panic(err)
	}
	// 检查连接
	err = mongoClient.Ping(context.TODO(), nil)
	if err != nil {
		panic(err)
	}
}

func GetMongoClient(uri string) *mongo.Client {
	once.Do(func() {
		InitEngine(uri)
	})

	return mongoClient
}

func NewMongoClient(uri string) (*mongo.Client, error) {
	cliOptions := options.Client().ApplyURI(uri)

	// 连接到 mongo
	client, err := mongo.Connect(context.TODO(), cliOptions)
	if err != nil {
		return nil, err
	}
	// 检查连接
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		return nil, err
	}

	MongoClient = client
	return client, nil
}

// 连接池模式
func ConnectToDB(uri, db string, timeout time.Duration, num uint64) (*mongo.Database, error) {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	o := options.Client().ApplyURI(uri)
	o.SetMaxPoolSize(num)
	client, err := mongo.Connect(ctx, o)
	if err != nil {
		return nil, err
	}

	return client.Database(db), nil
}

