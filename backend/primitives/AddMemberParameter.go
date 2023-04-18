package primitives

import (
	"time"
)

type AddMemberParameter struct {
	UserId           string      `bson:"user_id" json:"user_id"`
	Username         string      `bson:"username" json:"username"`
	Points           int         `bson:"points" json:"points"`
	Deadlines        []time.Time `bson:"deadlines" json:"deadlines"` // year month day hour minute seconds
	IsChecked        []bool      `bson:"is_checked" json:"is_checked"`
	IsReceivedPoints bool        `bson:"is_received_points" json:"is_received_points"`
}
