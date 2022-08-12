package userrepo

import (
	"context"
	"errors"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserRepo struct {
	Database string
	Client   *mongo.Client
}

const collectionname = "users"

func (ur *UserRepo) Insert(data map[string]interface{}) (map[string]interface{}, string, error) {
	collection := ur.Client.Database(ur.Database).Collection(collectionname)
	insertResult, err := collection.InsertOne(context.TODO(), data)
	if err != nil {
		fmt.Println(err)
		return nil, "702", err
	}
	result := make(map[string]interface{})
	err = collection.FindOne(context.TODO(), bson.M{"_id": insertResult.InsertedID}).Decode(&result)
	if err != nil {
		fmt.Println(err)
		return nil, "703", err
	}
	return result, "", nil
}

func (ur *UserRepo) IsNameNotExists(name string) (map[string]interface{}, string, error) {
	collection := ur.Client.Database(ur.Database).Collection(collectionname)
	result := make(map[string]interface{})
	err := collection.FindOne(context.TODO(), bson.M{"name": name}).Decode(&result)
	if err != nil {
		fmt.Println(err)
		if err == mongo.ErrNoDocuments {
			log.Println("No documents found")
			return nil, "", nil
		}
		return nil, "704", err
	}
	log.Println("Found a single document", result)
	return result, "705", errors.New("name already exists")
}
