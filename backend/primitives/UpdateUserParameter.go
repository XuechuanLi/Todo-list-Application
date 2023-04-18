package primitives

import (
	"go.mongodb.org/mongo-driver/bson"
	"time"
)

type UpdateUserParameter struct {
	Username    string    `bson:"username" json:"username"`
	Email       string    `bson:"email" json:"email"`
	Password    string    `bson:"password" json:"password"`
	TotalPoints int       `bson:"points" json:"points"`
	CreatedAt   time.Time `bson:"created_at" json:"created_at"`
	Filter      bson.D
}
