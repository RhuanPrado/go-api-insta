package user

import (
	"context"
	"go-api-insta/helpers/database"
	"go-api-insta/helpers/variable"
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

	var userExists *User

	filter := bson.D{{Key: "ID", Value: id}}

	err := repository.FindOne(context.TODO(), filter).Decode(userExists)
	if err != nil {
		println(err.Error())
		return nil, err
	}

	update := bson.D{{Key: "$set", Value: bson.D{{Key: "username", Value: username}}}}

	return repository.UpdateOne(context.TODO(), filter, update)
}

func (*UserRepository) GetUserByUsername(username string) (*User, error) {

	user := &User{}

	filter := bson.D{{Key: "username", Value: username}}

	err := repository.FindOne(context.TODO(), filter).Decode(user)
	if err != nil {
		println(err.Error())
		return nil, err
	}

	return user, nil
}
