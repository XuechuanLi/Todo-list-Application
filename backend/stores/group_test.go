package stores

import (
	"context"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"testing"
	"time"
	"todo-app/models"
	"todo-app/primitives"
	"todo-app/utils"
)

func TestGroupsDB_AddGroupToDb(t *testing.T) {
	db := NewGroupsDB(context.TODO(), utils.GetCollection("MONGO_DATABASE", "MONGO_GROUP_COLLECTION"))
	uidA, _ := primitive.ObjectIDFromHex("63b6e0d53ba748dda2eb25ad")
	uidB, _ := primitive.ObjectIDFromHex("63b6e0d53ba748dda2eb25ae")
	var newGroup = primitives.AddGroupParameter{
		Jackpot: 100, Title: "Learning Golang",
		Content: "A group that aims to help each other learn golang",
		Members: []models.Member{
			{
				UID:      uidA,
				Username: "A",
				Points:   50,
				Deadlines: []time.Time{
					time.Date(2022, time.December, 10, 10, 0, 0, 0, time.UTC),
					time.Date(2022, time.December, 11, 10, 0, 0, 0, time.UTC),
				},
				IsChecked: []bool{
					true,
					true,
				},
				IsReceivedPoints: false,
			},
			{
				UID:      uidB,
				Username: "B",
				Points:   50,
				Deadlines: []time.Time{
					time.Date(2022, time.December, 10, 10, 0, 0, 0, time.UTC),
					time.Date(2022, time.December, 11, 10, 0, 0, 0, time.UTC),
				},
				IsChecked: []bool{
					false,
					false,
				},
				IsReceivedPoints: false,
			},
		},
	}
	result := db.AddGroupToDb(newGroup)
	if result == nil {
		t.Errorf("insert group failed")
	}
}
