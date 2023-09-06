package user

import (
	"context"
	"errors"
	"go-api-insta/helpers/database"
	"go-api-insta/helpers/variable"
	"go-api-insta/libs/logger"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var repository = database.Dbconnect().Database(variable.GetEnvVariable("DATABASE_NAME")).Collection("users")

type UserRepository struct{}

func (*UserRepository) CreateUser(newUser *User) (*mongo.InsertOneResult, error) {
	newUser.ID = primitive.NewObjectID()
	newUser.CreatedAt = time.Now()
	newUser.UpdatedAt = time.Now()
	return repository.InsertOne(context.TODO(), *newUser)
}

func (*UserRepository) UpdatedUserName(id string, username string) (*mongo.UpdateResult, error) {

	userExists := &User{}

	docId, _ := primitive.ObjectIDFromHex(id)

	filter := bson.M{"_id": docId}

	err := repository.FindOne(context.TODO(), filter).Decode(userExists)
	if err != nil {
		println(err.Error())
		return nil, err
	}
	if userExists.Username == "" {
		return nil, errors.New("user dont exists")
	}

	update := bson.D{{Key: "$set", Value: bson.D{{Key: "username", Value: username}}}}

	return repository.UpdateOne(context.TODO(), filter, update)
}

func (*UserRepository) GetUserByUsername(username string) (*User, error) {

	user := &User{}

	filter := bson.D{{Key: "username", Value: username}}

	err := repository.FindOne(context.TODO(), filter).Decode(user)
	if err != nil {
		return nil, err
	}

	if user.Username == "" {
		return nil, nil
	}

	return user, nil
}

func (*UserRepository) GetUserById(id string) (*User, error) {

	user := &User{}

	docId, _ := primitive.ObjectIDFromHex(id)

	filter := bson.M{"_id": docId}

	err := repository.FindOne(context.TODO(), filter).Decode(user)
	if err != nil {
		return nil, err
	}

	if user.Username == "" {
		return nil, nil
	}

	return user, nil
}

func (*UserRepository) UpdatedFriends(id string, friends []string) (*mongo.UpdateResult, error) {

	userExists := &User{}

	docId, _ := primitive.ObjectIDFromHex(id)

	filter := bson.M{"_id": docId}

	err := repository.FindOne(context.TODO(), filter).Decode(userExists)
	if err != nil {
		println(err.Error())
		return nil, err
	}
	if userExists.Username == "" {
		return nil, errors.New("user dont exists")
	}

	update := bson.D{{Key: "$set", Value: bson.D{{Key: "friends", Value: friends}}}}

	return repository.UpdateOne(context.TODO(), filter, update)
}

func (*UserRepository) ListUsers(id string) ([]User, error) {

	userExists := &User{}

	filter := bson.M{"_id": bson.M{"$ne": id}}

	err := repository.FindOne(context.TODO(), filter).Decode(userExists)
	if err != nil {
		println(err.Error())
		return nil, err
	}
	if userExists.Username == "" {
		return nil, errors.New("user dont exists")
	}

	cursor, err := repository.Find(context.TODO(), filter)
	if err != nil {
		return nil, err
	}
	users := []User{}
	for cursor.Next(context.TODO()) {
		//Create a value into which the single document can be decoded
		var elem User
		err := cursor.Decode(&elem)
		if err != nil {
			logger.Production.Info("error decode post element array")
		}

		users = append(users, elem)
	}
	return users, nil
}
