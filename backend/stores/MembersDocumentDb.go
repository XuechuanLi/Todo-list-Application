package stores

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
	. "todo-app/models"
	. "todo-app/primitives"
	"todo-app/utils"
)

var (
	memberCollection = utils.GetCollection("MONGO_DATABASE", "MONGO_MEMBER_COLLECTION")
)

func GetMemberFromDb(filter bson.D) (resultMember *Member, err error) {
	err = memberCollection.FindOne(context.TODO(), filter).Decode(&resultMember)

	if err != nil {
		if err == mongo.ErrNoDocuments {
			return
		}
		panic(err)
	}
	return resultMember, nil
}

func BulkGetMembersFromDb(filter bson.D) (resultMembers []Member) {
	bulkGetResponse, err := memberCollection.Find(context.TODO(), filter)
	if err != nil {
		log.Fatal(err)
	}

	defer func(bulkGetResponse *mongo.Cursor, ctx context.Context) {
		err := bulkGetResponse.Close(ctx)
		if err != nil {

		}
	}(bulkGetResponse, context.TODO())

	for bulkGetResponse.Next(context.TODO()) {

		var result Member
		err := bulkGetResponse.Decode(&result)
		if err != nil {
			log.Fatal(err)
		}
		resultMembers = append(resultMembers, result)
	}
	if err := bulkGetResponse.Err(); err != nil {
		log.Fatal(err)
	}
	return resultMembers
}

func AddMemberToDb(addMemberParameter AddMemberParameter) *mongo.InsertOneResult {
	uid, err := primitive.ObjectIDFromHex(addMemberParameter.UserId)
	if err != nil {
		panic(err)
	}
	newMember := Member{
		UID:              uid,
		Username:         addMemberParameter.Username,
		Points:           addMemberParameter.Points,
		Deadlines:        addMemberParameter.Deadlines,
		IsChecked:        addMemberParameter.IsChecked,
		IsReceivedPoints: addMemberParameter.IsReceivedPoints,
	}
	response, err := memberCollection.InsertOne(context.Background(), newMember)
	if err != nil {
		panic(err)
	}
	return response
}

func BulkAddMembersToDb(bulkAddMembersParameter BulkAddMembersParameters) *mongo.InsertManyResult {
	var newMembers []interface{}
	for index := 0; index < len(bulkAddMembersParameter.Parameters); index++ {
		addMemberParameter := bulkAddMembersParameter.Parameters[index]
		newMember := Member{
			UID:              primitive.NewObjectID(),
			Username:         addMemberParameter.Username,
			Points:           addMemberParameter.Points,
			Deadlines:        addMemberParameter.Deadlines,
			IsChecked:        addMemberParameter.IsChecked,
			IsReceivedPoints: addMemberParameter.IsReceivedPoints,
		}
		newMembers = append(newMembers, newMember)
	}
	response, err := memberCollection.InsertMany(context.TODO(), newMembers)
	if err != nil {
		panic(err)
	}
	return response
}

func UpdateMemberToDb(updateMemberParameter UpdateMemberParameter) *mongo.UpdateResult {
	getMemberResponse, getErr := GetMemberFromDb(updateMemberParameter.Filter)
	if getErr != nil {
		panic("Error: The original user do not exist.")
	}
	newMember := Member{
		UID:              getMemberResponse.UID,
		Username:         updateMemberParameter.Username,
		Points:           updateMemberParameter.Points,
		Deadlines:        updateMemberParameter.Deadlines,
		IsChecked:        updateMemberParameter.IsChecked,
		IsReceivedPoints: updateMemberParameter.IsReceivedPoints,
	}

	updateMemberResponse, err := memberCollection.ReplaceOne(context.TODO(), updateMemberParameter.Filter, newMember)
	if err != nil {
		panic(err)
	}

	return updateMemberResponse
}

func DeleteMemberFromDb(filter bson.D) *mongo.DeleteResult {
	response, err := memberCollection.DeleteOne(context.TODO(), filter)
	if err != nil {
		panic(err)
	}
	return response
}
