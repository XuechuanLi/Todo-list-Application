package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type Member struct {
	UID              primitive.ObjectID `bson:"_id" json:"uid"`
	Username         string             `bson:"username" json:"username"`
	Points           int                `bson:"points" json:"points"`
	Deadlines        []time.Time        `bson:"deadlines" json:"deadlines"` // year month day hour minute seconds
	IsChecked        []bool             `bson:"is_checked" json:"is_checked"`
	IsReceivedPoints bool               `bson:"is_received_points" json:"is_received_points"`
}

type Group struct {
	ID      primitive.ObjectID `bson:"_id" json:"gid"`
	Jackpot int                `bson:"jackpot" json:"jackpot"`
	Title   string             `bson:"title" json:"title"`
	Content string             `bson:"content" json:"content"`
	Members []Member           `bson:"members" json:"members"`
}
