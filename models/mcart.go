package models

import (
	"context"
	"log"
	"softwarecart/helper"
	"time"

	"go.mongodb.org/mongo-driver/mongo/options"
)

type Shopingcart struct {
	Id         string    `json:"_id,omitempty" bson:"_id,omitempty"`
	Kuantitas  int       `bson:"kuantitas" json:"kuantitas"`
	Produkid   string    `bson:"produkid" json:"produkid"`
	Produkname string    `bson:"produkname" json:"produkname"`
	Created_at time.Time `json:"created_at" bson:"created_at"`
}

func FindCart(filter interface{}, limit int, offset int) ([]Shopingcart, error) {
	collection := helper.ConnectDB("carts")

	findOptions := options.Find()
	if limit != 0 {
		findOptions.SetLimit(int64(limit))
	}
	findOptions.SetSkip(int64(offset))
	ctx, cancel := helper.ContextTimeout(0)
	defer cancel()
	cur, findError := collection.Find(context.TODO(), filter, findOptions)
	returnData := []Shopingcart{}
	if findError != nil {
		return returnData, findError
	}

	//Map result to slice
	for cur.Next(context.TODO()) {
		t := Shopingcart{}

		err := cur.Decode(&t)
		if err != nil {
			log.Fatal(err)
		}
		returnData = append(returnData, t)
	}
	defer cur.Close(ctx)

	return returnData, nil
}

func InserData(insert interface{}) map[string]interface{} {
	collection := helper.ConnectDB("carts")
	_, err := collection.InsertOne(context.TODO(), insert)
	returnData := map[string]interface{}{
		"status": 0,
	}
	if err != nil {
		returnData["status"] = 0
	} else {
		returnData["status"] = 1
	}

	return returnData
}

func UpdateOne(filter interface{}, data interface{}) bool {
	collection := helper.ConnectDB("carts")
	_, err := collection.UpdateOne(context.TODO(), filter, data)
	if err != nil {
		return false
	}

	return true
}

func DeleteOne(filter interface{}) bool {
	collection := helper.ConnectDB("carts")
	_, err := collection.DeleteOne(context.TODO(), filter)
	if err != nil {
		return false
	}
	return true
}

func FindOne(filter interface{}) (Shopingcart, error) {
	collection := helper.ConnectDB("carts")
	var tempData Shopingcart
	findError := collection.FindOne(context.TODO(), filter).Decode(&tempData)
	if findError != nil {
		return tempData, findError
	}
	log.Println("masuk sini : ", tempData)

	return tempData, nil
}
