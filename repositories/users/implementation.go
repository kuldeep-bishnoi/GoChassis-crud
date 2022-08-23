package userrepo

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"strconv"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type UserRepo struct {
	Database string
	Client   *mongo.Client
}

const collectionname = "users"

func init() { log.SetFlags(log.Lshortfile | log.LstdFlags) }

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
	return result, "716", nil
}

func (ur *UserRepo) IsNameNotExists(name string) (map[string]interface{}, string, error) {
	collection := ur.Client.Database(ur.Database).Collection(collectionname)
	result := make(map[string]interface{})
	err := collection.FindOne(context.TODO(), bson.M{"name": name}).Decode(&result)
	if err != nil {
		fmt.Println(err)
		if err == mongo.ErrNoDocuments {
			log.Println("No documents found")
			return nil, "713", nil
		}
		return nil, "717", err
	}
	log.Println("Found a single document", result)
	return nil, "718", errors.New("name already exists")
}

func (ur *UserRepo) GetAllUsersInput(filters string, page string, limit string) ([]map[string]interface{}, string, error, int) {
	collection := ur.Client.Database(ur.Database).Collection(collectionname)
	p := make(map[string]interface{})
	err := json.Unmarshal([]byte(filters), &p)
	if err != nil {
		var result []map[string]interface{}
		totalcount, err := collection.CountDocuments(context.TODO(), bson.M{})
		if err != nil {
			log.Println(err)
			return nil, "720", err, 0 //fix later error code
		}
		cur, err := collection.Find(context.TODO(), bson.M{})
		if err != nil {
			if err == mongo.ErrNoDocuments {
				log.Println("no documents found")
				return nil, "721", err, 0
			}
			return nil, "722", err, 0
		}
		for cur.Next(context.TODO()) {
			var doc map[string]interface{}
			err := cur.Decode(&doc)
			if err != nil {
				log.Println(err.Error())
				return nil, "723", err, 0 //fix later error code
			}
			result = append(result, doc)
		}
		cur.Close(context.TODO()) // close the cursor once stream of documents has exhausted
		fmt.Println("total count1:", totalcount)
		return result, "", nil, int(totalcount)
	}
	if page == "" {
		page = "1"
	}
	if limit == "" {
		limit = "10"
	}
	var results []map[string]interface{}
	pno, err := strconv.ParseInt(page, 10, 64)
	if err != nil {
		fmt.Println(err)
		return nil, "711", err, 0
	}
	plim, err := strconv.ParseInt(limit, 10, 64)
	if err != nil {
		fmt.Println(err)
		return nil, "712", err, 0
	}
	skip := (pno - 1) * plim
	opts := options.FindOptions{Skip: &skip, Limit: &plim}
	totalcount, err := collection.CountDocuments(context.TODO(), p)
	if err != nil {
		fmt.Println(err)
		return nil, "721", err, 0
	}
	if totalcount == 0 {
		return nil, "721", errors.New("totalcount is empty"), 0
	}
	cur, err := collection.Find(context.TODO(), p, &opts)
	fmt.Println(err)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			log.Println("no documents found")
			return nil, "721", err, 0
		}
		return nil, "722", err, 0
	}
	for cur.Next(context.TODO()) {
		var doc map[string]interface{}
		err := cur.Decode(&doc)
		if err != nil {
			log.Println(err.Error())
			return nil, "723", err, 0 //fix later error code
		}
		results = append(results, doc)
	}
	cur.Close(context.TODO()) // close the cursor once stream of documents has exhausted
	// common.ResponseHandler("710", "en", int(totalcount), results)
	return results, "", err, int(totalcount)
}

func (ur *UserRepo) Delete(id string) (map[string]interface{}, string, error) {
	collection := ur.Client.Database(ur.Database).Collection(collectionname)
	Id, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, "725", errors.New("invalid id")
	}
	result := make(map[string]interface{})
	err = collection.FindOneAndDelete(context.TODO(), bson.M{"_id": Id}).Decode(&result)
	if err != nil {
		return nil, "724", err
	}
	return result, "", nil

}

func (ur *UserRepo) Getbyid(id string) (map[string]interface{}, string, error) {
	collection := ur.Client.Database(ur.Database).Collection(collectionname)
	Id, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		log.Println(err.Error())
		return nil, "726", err
	}
	result := make(map[string]interface{})
	err = collection.FindOne(context.TODO(), bson.M{"_id": Id}).Decode(&result)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			log.Println("no documents found line 171")
			return result, "713", errors.New("no Documents Found")
		}
		log.Println("Internal Server Error")
		return nil, "727", errors.New("internal Server Error")
	}
	return result, "", nil
}

func (ur *UserRepo) Update(id string, data map[string]interface{}) (map[string]interface{}, string, error) {
	collection := ur.Client.Database(ur.Database).Collection(collectionname)
	update := make(map[string]interface{})
	update["$set"] = data
	Id, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		log.Println(err)
		return nil, "706", err
	}
	result := make(map[string]interface{})
	err = collection.FindOneAndUpdate(context.TODO(), bson.M{"_id": Id}, update).Decode(&result)
	if err != nil {
		log.Println(err.Error())
		return nil, "715", err
	}
	result, _, err = ur.Getbyid(string(id))
	if err != nil {
		return nil, "728", err
	}
	return result, "", nil
}
