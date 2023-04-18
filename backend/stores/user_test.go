package stores

import (
	"context"
	"testing"
	"time"
	"todo-app/primitives"
	"todo-app/utils"
)

func TestUsersDB_AddUserToDb(t *testing.T) {
	db := NewUsersDB(context.TODO(), utils.GetCollection("MONGO_DATABASE", "MONGO_USER_COLLECTION"))
	var testUser = primitives.AddUserParameter{
		Username:    "D",
		Email:       "D@example.com",
		Password:    "SDF",
		TotalPoints: 1000,
		CreatedAt:   time.Now(),
	}
	result := db.AddUserToDb(testUser)
	if result == nil {
		t.Errorf("insert failed")
	}
}
