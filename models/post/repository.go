package post

import (
	"context"
	"go-api-insta/helpers/database"
	"go-api-insta/helpers/variable"
	"go-api-insta/libs/logger"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var repository = database.Dbconnect().Database(variable.GetEnvVariable("DATABASE_NAME")).Collection("posts")

type PostRepository struct{}

func (*PostRepository) CreatePost(userId string, newPost *Post) (*mongo.InsertOneResult, error) {
	newPost.ID = primitive.NewObjectID()
	newPost.CreatedAt = time.Now()
	newPost.UpdatedAt = time.Now()
	return repository.InsertOne(context.TODO(), *newPost)
}

func (*PostRepository) FindAllPostByUser(id string) ([]Post, error) {
	filter := bson.D{{Key: "userId", Value: id}}
	cursor, err := repository.Find(context.TODO(), filter)
	if err != nil {
		return nil, err
	}
	postsUser := []Post{}

	for cursor.Next(context.TODO()) {
		//Create a value into which the single document can be decoded
		var elem Post
		err := cursor.Decode(&elem)
		if err != nil {
			logger.Production.Info("error decode post element array")
		}

		postsUser = append(postsUser, elem)

	}

	return postsUser, nil
}

func (*PostRepository) FindAllPostFriends(friends []string) ([]Post, error) {

	filter := bson.M{"userId": bson.M{"$in": friends}}
	options := options.Find().SetSort(bson.D{{Key: "created_at", Value: 1}})
	cursor, err := repository.Find(context.TODO(), filter, options)
	if err != nil {
		return nil, err
	}
	postsUser := []Post{}

	for cursor.Next(context.TODO()) {
		//Create a value into which the single document can be decoded
		var elem Post
		err := cursor.Decode(&elem)
		if err != nil {
			logger.Production.Info("error decode post element array")
		}

		postsUser = append(postsUser, elem)

	}

	return postsUser, nil
}
