package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type User struct {
	ID          primitive.ObjectID `bson:"_id" json:"uid"`
	Username    string             `bson:"username" json:"username"`
	Email       string             `bson:"email" json:"email"`
	Password    string             `bson:"password" json:"password"`
	TotalPoints int                `bson:"points" json:"points"`
	CreatedAt   time.Time          `bson:"created_at" json:"created_at"`
}
