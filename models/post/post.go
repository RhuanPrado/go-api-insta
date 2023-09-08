package post

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Post struct {
	ID          primitive.ObjectID `json:"id" bson:"_id"`
	CreatedAt   time.Time          `json:"created_at" bson:"created_at"`
	UpdatedAt   time.Time          `json:"updated_at" bson:"updated_at"`
	Description string             `json:"description" bson:"description"`
	File        []byte             `json:"file" bson:"file"`
	UserId      string             `json:"userId" bson:"userId"`
}
