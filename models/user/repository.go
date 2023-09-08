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
		return nil, errors.New("users does not exists")
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

func (*UserRepository) ListUsers(id string, friends []string) ([]User, error) {

	usersFilter := append(friends, id)
	objIds := listIdToObejectId(usersFilter)
	filter := bson.M{"_id": bson.M{"$nin": objIds}}

	println(filter)

	cursor, err := repository.Find(context.TODO(), filter)
	if err != nil {
		return nil, err
	}
	users := []User{}
	for cursor.Next(context.TODO()) {

		var elem User
		err := cursor.Decode(&elem)
		if err != nil {
			logger.Production.Info("error decode post element array")
		}

		users = append(users, elem)
	}
	return users, nil
}

func (*UserRepository) ListUsersFriends(id string, friends []string) ([]User, error) {

	objIds := listIdToObejectId(friends)

	filter := bson.M{"_id": bson.M{"$in": objIds}}

	cursor, err := repository.Find(context.TODO(), filter)
	if err != nil {
		return nil, err
	}
	users := []User{}
	for cursor.Next(context.TODO()) {

		var elem User
		err := cursor.Decode(&elem)
		if err != nil {
			logger.Production.Info("error decode post element array")
		}

		users = append(users, elem)
	}
	return users, nil
}

func listIdToObejectId(list []string) []primitive.ObjectID {
	listObjId := []primitive.ObjectID{}
	for _, ele := range list {
		objId, err := primitive.ObjectIDFromHex(ele)
		if err != nil {
			logger.Production.Info("error objectId")
		}
		listObjId = append(listObjId, objId)
	}
	return listObjId
}
