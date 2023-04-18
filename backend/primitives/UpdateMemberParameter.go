package primitives

import (
	"go.mongodb.org/mongo-driver/bson"
	"time"
)

type UpdateMemberParameter struct {
	Username         string      `bson:"username" json:"username"`
	Points           int         `bson:"points" json:"points"`
	Deadlines        []time.Time `bson:"deadlines" json:"deadlines"` // year month day hour minute seconds
	IsChecked        []bool      `bson:"is_checked" json:"is_checked"`
	IsReceivedPoints bool        `bson:"is_received_points" json:"is_received_points"`
	Filter           bson.D
}
