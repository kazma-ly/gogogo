package llog

import (
	"context"
	"fmt"
	"log"

	"github.com/mongodb/mongo-go-driver/mongo"
)

var (
	collection *mongo.Collection
)

// ProcessLog 日志处理
func ProcessLog(val interface{}) {
	fmt.Println(val)
	_, err := GetCollection().InsertOne(context.Background(), val)
	if err != nil {
		log.Printf("save error: %s", err.Error())
	}
}

func GetCollection() *mongo.Collection {
	if collection == nil {
		client, err := mongo.NewClient("mongodb://admin:Kazma123456.@kazma.life:27017")
		if err != nil {
			log.Fatal(err)
		}
		err = client.Connect(context.TODO())
		if err != nil {
			log.Fatal(err)
		}
		collection = client.Database("Logs").Collection("LifeLog")
	}
	return collection
}
