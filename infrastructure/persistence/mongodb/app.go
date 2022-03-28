package mongodb

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"cc/domain/entity"
)

type AppMongoHandler struct {
	client     *mongo.Client
	db         *mongo.Database
	collection *mongo.Collection
}

func NewAppMongoHandler(client *mongo.Client, database, collection string) *AppMongoHandler {
	db := client.Database(database)
	col := db.Collection(collection)
	return &AppMongoHandler{
		client:     client,
		db:         db,
		collection: col,
	}
}

func (handler *AppMongoHandler) Save(app entity.App) error {
	id, err := handler.findLastID()
	if err != nil {
		return err
	}
	app.ID = id + 1
	_, err = handler.collection.InsertOne(context.TODO(), app)
	return err
}

func (handler *AppMongoHandler) FindAll() ([]*entity.App, error) {
	var apps []*entity.App
	opts := options.Find()
	ctx := context.TODO()
	cur, err := handler.collection.Find(ctx, bson.D{{}}, opts)
	if err != nil {
		return nil, err
	}
	defer cur.Close(ctx)

	// 查找多个文档返回一个光标
	// 遍历游标允许我们一次解码一个文档
	for cur.Next(ctx) {
		// 创建一个值，将单个文档解码为该值
		var app *entity.App
		err = cur.Decode(&app)
		if err != nil {
			return nil, err
		}
		apps = append(apps, app)
	}

	return apps, nil
}

func (handler *AppMongoHandler) FindOne(id uint64) (app *entity.App, err error) {
	filter := bson.D{{"id", id}}

	err = handler.collection.FindOne(context.TODO(), filter).Decode(&app)
	if err != nil {
		return nil, err
	}
	return app, nil
}

func (handler *AppMongoHandler) Delete(id uint64) error {
	res, err := handler.collection.DeleteOne(context.TODO(), bson.D{{"id", id}})
	fmt.Println(res)
	return err
}

// 更新一条数据
func (handler *AppMongoHandler) Update(app entity.App) error {
	filter := bson.M{"id": app.ID}
	data := make(map[string]interface{})
	if app.Name != "" {
		data["name"] = app.Name
	}
	if app.Version != "" {
		data["version"] = app.Version
	}
	if len(data) == 0 {
		return nil
	}
	data["updated_at"] = time.Now()
	update := bson.M{"$set": bson.M(data)}
	_, err := handler.collection.UpdateOne(context.TODO(), filter, update)
	return err
}

// 更新多条数据
func (handler *AppMongoHandler) UpdateMany() error {
	return nil
}

// 批量删除
func (handler *AppMongoHandler) BatchDelete() error {
	return nil
}

// 查找最新ID
func (handler *AppMongoHandler) findLastID() (uint64, error) {
	cur, err := handler.collection.Find(context.TODO(), bson.D{{}}, options.Find().SetLimit(1), options.Find().SetSort(bson.M{"id": -1}))
	if err != nil {
		return 0, err
	}
	defer cur.Close(context.TODO())
	if cur.Err() != nil {
		return 0, cur.Err()
	}

	var app entity.App
	for cur.Next(context.TODO()) {
		err = cur.Decode(&app)
		if err != nil {
			return 0, err
		}
	}

	return app.ID, nil
}