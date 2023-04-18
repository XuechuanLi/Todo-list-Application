package primitives

import "time"

type AddUserParameter struct {
	Username    string    `bson:"username" json:"username"`
	Email       string    `bson:"email" json:"email"`
	Password    string    `bson:"password" json:"password"`
	TotalPoints int       `bson:"points" json:"points"`
	CreatedAt   time.Time `bson:"created_at" json:"created_at"`
}
