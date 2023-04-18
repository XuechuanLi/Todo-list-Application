package unitTest

//
//import (
//	"go.mongodb.org/mongo-driver/bson"
//	"log"
//	"testing"
//	"time"
//	. "todo-app/primitives"
//	. "todo-app/stores"
//)
//
//func TestAddUser(t *testing.T) {
//	var testUser = AddUserParameter{
//		Username:    "A",
//		Email:       "A@example.com",
//		Password:    "xxxx",
//		TotalPoints: 1000,
//		CreatedAt:   time.Date(2020, time.December, 10, 23, 1, 10, 0, time.UTC),
//	}
//	t.Run("addUserTest", func(t *testing.T) {
//		response := AddUserToDb(testUser)
//		log.Fatal(response)
//	})
//}
//
//func TestBulkAddUsers(t *testing.T) {
//	var testUser1 = AddUserParameter{
//		Username:    "B",
//		Email:       "B@example.com",
//		Password:    "ASBF",
//		TotalPoints: 100,
//		CreatedAt:   time.Date(2020, time.December, 10, 23, 1, 10, 0, time.UTC),
//	}
//	var testUser2 = AddUserParameter{
//		Username:    "C",
//		Email:       "C@example.com",
//		Password:    "ASBF",
//		TotalPoints: 100,
//		CreatedAt:   time.Date(2020, time.December, 10, 23, 1, 10, 0, time.UTC),
//	}
//	var testUsers = BulkAddUsersParameters{
//		Parameters: []AddUserParameter{testUser1, testUser2},
//	}
//
//	t.Run("bulkAddUsersTest", func(t *testing.T) {
//		BulkAddUsersToDb(testUsers)
//	})
//}
//
//func TestGetUser(t *testing.T) {
//	filter := bson.D{{"username", "A"}}
//	t.Run("getUserTest", func(t *testing.T) {
//		response, err := GetUserFromDb(filter)
//		if err != nil {
//			panic(err)
//		}
//		print(response)
//	})
//}
//
//func TestBulkGetUsers(t *testing.T) {
//	filter := bson.D{{"password", "ASBF"}}
//	t.Run("bulkGetUsersTest", func(t *testing.T) {
//		response := BulkGetUsersFromDb(filter)
//		print(response)
//	})
//}
//
//func TestDeleteUser(t *testing.T) {
//	filter := bson.D{{"username", "A"}}
//	t.Run("deleteUserTest", func(t *testing.T) {
//		response := DeleteUserFromDb(filter)
//		print(response)
//	})
//}
//
//func TestUpdateUser(t *testing.T) {
//	filter := bson.D{{"username", "C"}}
//	updateUserParameter := UpdateUserParameter{
//		Username:    "D",
//		Email:       "C@example.com",
//		Password:    "ASBF",
//		TotalPoints: 100,
//		CreatedAt:   time.Date(2020, time.December, 10, 23, 1, 10, 0, time.UTC),
//		Filter:      filter,
//	}
//	t.Run("updateUserTest", func(t *testing.T) {
//		response := UpdateUserToDb(updateUserParameter)
//		print(response)
//	})
//}
//
//func TestAddMember(t *testing.T) {
//	var testMember = AddMemberParameter{
//		UserId:           "63ae4ff2db08212ae6f8d56f",
//		Username:         "B",
//		Points:           10,
//		Deadlines:        []time.Time{time.Now(), time.Now()},
//		IsChecked:        []bool{true, true},
//		IsReceivedPoints: false,
//	}
//	t.Run("addMemberTest", func(t *testing.T) {
//		response := AddMemberToDb(testMember)
//		print(response)
//	})
//}
//
//func TestBulkAddMembers(t *testing.T) {
//	var testMember1 = AddMemberParameter{
//		UserId:           "63ae4ff2db08212ae6f8d56f",
//		Username:         "B",
//		Points:           10,
//		Deadlines:        []time.Time{time.Now(), time.Now()},
//		IsChecked:        []bool{true, true},
//		IsReceivedPoints: false,
//	}
//	var testMember2 = AddMemberParameter{
//		UserId:           "63ae4ff2db08212ae6f8d570",
//		Username:         "D",
//		Points:           10,
//		Deadlines:        []time.Time{time.Now(), time.Now()},
//		IsChecked:        []bool{true, true},
//		IsReceivedPoints: false,
//	}
//	var testMembers = BulkAddMembersParameters{
//		Parameters: []AddMemberParameter{testMember1, testMember2},
//	}
//
//	t.Run("bulkAddMembersTest", func(t *testing.T) {
//		BulkAddMembersToDb(testMembers)
//	})
//}
//
//func TestGetMember(t *testing.T) {
//	filter := bson.D{{"user_id", "63ae4ff2db08212ae6f8d570"}}
//	t.Run("getMemberTest", func(t *testing.T) {
//		response, err := GetMemberFromDb(filter)
//		if err != nil {
//			panic(err)
//		}
//		print(response)
//	})
//}
//
//func TestBulkGetMembers(t *testing.T) {
//	filter := bson.D{{"username", "B"}}
//	t.Run("bulkGetMembersTest", func(t *testing.T) {
//		response := BulkGetMembersFromDb(filter)
//		print(response)
//	})
//}
//
//func TestDeleteMember(t *testing.T) {
//	filter := bson.D{{"username", "B"}}
//	t.Run("deleteMemberTest", func(t *testing.T) {
//		response := DeleteMemberFromDb(filter)
//		print(response)
//	})
//}
//
//func TestUpdateMember(t *testing.T) {
//	filter := bson.D{{"username", "D"}}
//	updateMemberParameter := UpdateMemberParameter{
//		Username:         "E",
//		Points:           10,
//		Deadlines:        []time.Time{time.Now(), time.Now()},
//		IsChecked:        []bool{true, true},
//		IsReceivedPoints: false,
//		Filter:           filter,
//	}
//	t.Run("updateMemberTest", func(t *testing.T) {
//		response := UpdateMemberToDb(updateMemberParameter)
//		print(response)
//	})
//}
