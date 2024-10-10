package services

import (
	"context"
	"log"
	"time"

	"github.com/cafedex-backend/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserGuide models.Guide

var client *mongo.Client

func New(mongo *mongo.Client) UserGuide {
	client = mongo

	return UserGuide{}
}

func returnCollectionPointer(collection string) *mongo.Collection {
	return client.Database("cafedex_db").Collection(collection)
}

func GetAllGuides() ([]UserGuide, error) {
	collection := returnCollectionPointer("guides")
	var guides []UserGuide

	cursor, err := collection.Find(context.TODO(), bson.D{})
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	defer cursor.Close(context.Background())

	for cursor.Next(context.Background()) {
		var guide UserGuide
		cursor.Decode(&guide)
		guides = append(guides, guide)
	}

	return guides, nil
}

func GetGuideById(id string) (UserGuide, error) {
	collection := returnCollectionPointer("cafedex-guides")
	var todo UserGuide

	mongoID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return UserGuide{}, err
	}

	err = collection.FindOne(context.Background(), bson.M{"_id": mongoID}).Decode(&todo)
	if err != nil {
		log.Println(err)
		return UserGuide{}, err
	}

	return todo, nil
}

// func (t *Guide) GetGuideByAuthor(author string) ([]Guide, error) {
// 	collection := returnCollectionPointer("guides")
// 	var guides []Guide 
// 	cursor, err :=collection.Find(context.TODO(), bson.D{})
// 	if err != nil {
// 		log.Fatal(err)
// 		return nil, err
// 	}
// 	defer cursor.Close(context.Background())
// 	for cursor.Next(context.Background()){
// 		var guide Guide
// 		cursor.Decode(&author)
// 		guides.append(guides, guide)
// 	}
// 	err = collection.FindOne(context.Background(), bson.M{"_id": mongoID}).Decode(&todo)
// 	if err != nil {
// 		log.Println(err)
// 		return Guides{}, err
// 	}
// 	return guides, nil
// }

func UpdateGuide(id string, entry UserGuide){
	// TODO: Create the update schema for the backend
	// collection := returnCollectionPointer("cafedex-guides")
	// mongoID, err := primitive.ObjectIDFromHex(id)
	// if err != nil {
	// 	return nil, err
	// }

	// update := bson.D{
	// 	{"set", bson.D{


	// 	}},
	// }

}

func CreateGuide(entry UserGuide) error{
	collection := returnCollectionPointer("cafedex-guides")

	_, err := collection.InsertOne(context.TODO(), UserGuide{
		Author:     entry.Author,
		Title: 			entry.Title,
		CreatedAt: 	time.Now(),
		UpdatedAt: 	time.Now(),
	})
	if err != nil {
		log.Println("Error:", err)
		return err
	}

	return nil
}


func DeleteGuide(id string) error{
	collection := returnCollectionPointer("cafedex-guides")
	mongoID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		log.Println(err)
		return err
	}

	_, err = collection.DeleteOne(
		context.Background(),
		bson.M{"_id": mongoID},
	)
	if err != nil {
		log.Println(err)
		return err
	}

	return nil
}