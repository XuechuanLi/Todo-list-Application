package primitives

import (
	. "todo-app/models"
)

type AddGroupParameter struct {
	Jackpot int      `bson:"jackpot" json:"jackpot"`
	Title   string   `bson:"title" json:"title"`
	Content string   `bson:"content" json:"content"`
	Members []Member `bson:"members" json:"members"`
}
