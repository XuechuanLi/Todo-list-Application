package primitives

import (
	"go.mongodb.org/mongo-driver/bson"
	. "todo-app/models"
)

type UpdateGroupParameter struct {
	Jackpot int      `bson:"jackpot" json:"jackpot"`
	Title   string   `bson:"title" json:"title"`
	Content string   `bson:"content" json:"content"`
	Members []Member `bson:"members" json:"members"`
	Filter  bson.D
}
