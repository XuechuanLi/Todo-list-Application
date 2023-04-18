package stores

import (
	"context"
	"errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
	. "todo-app/models"
	. "todo-app/primitives"
)

type GroupsDB struct {
	collection *mongo.Collection
	ctx        context.Context
}

func NewGroupsDB(ctx context.Context, collection *mongo.Collection) *GroupsDB {
	return &GroupsDB{
		collection: collection,
		ctx:        ctx,
	}
}

//var (
//	groupCollection = utils.GetCollection("MONGO_DATABASE", "MONGO_GROUP_COLLECTION")
//)

func (g *GroupsDB) GetGroupFromDb(filter bson.D) (resultGroup *Group, err error) {
	err = g.collection.FindOne(g.ctx, filter).Decode(&resultGroup)

	if err != nil {
		if err == mongo.ErrNoDocuments {
			return
		}
		panic(err)
	}
	return resultGroup, nil
}

func (g *GroupsDB) BulkGetGroupsFromDb(filter bson.D) (resultGroups []Group, err error) {
	bulkGetResponse, err := g.collection.Find(g.ctx, filter)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	defer func(bulkGetResponse *mongo.Cursor, ctx context.Context) {
		err := bulkGetResponse.Close(ctx)
		if err != nil {

		}
	}(bulkGetResponse, context.TODO())

	for bulkGetResponse.Next(context.TODO()) {

		var result Group
		err := bulkGetResponse.Decode(&result)
		if err != nil {
			log.Fatal(err)
			return nil, err
		}
		resultGroups = append(resultGroups, result)
	}
	if err := bulkGetResponse.Err(); err != nil {
		log.Fatal(err)
		return nil, err
	}
	return resultGroups, nil
}

func (g *GroupsDB) AddGroupToDb(newGroupParameter AddGroupParameter) (*mongo.InsertOneResult, error) {
	newGroup := Group{
		ID:      primitive.NewObjectID(),
		Jackpot: newGroupParameter.Jackpot,
		Title:   newGroupParameter.Title,
		Content: newGroupParameter.Content,
		Members: newGroupParameter.Members,
	}
	response, err := g.collection.InsertOne(g.ctx, newGroup)
	if err != nil {
		return nil, err
	}
	return response, nil
}

func (g *GroupsDB) BulkAddGroupToDb(bulkAddGroupsParameter BulkAddGroupsParameters) *mongo.InsertManyResult {
	var newGroups []interface{}
	for index := 0; index < len(bulkAddGroupsParameter.Parameters); index++ {
		addGroupParameter := bulkAddGroupsParameter.Parameters[index]
		newGroup := Group{
			ID:      primitive.NewObjectID(),
			Jackpot: addGroupParameter.Jackpot,
			Title:   addGroupParameter.Title,
			Content: addGroupParameter.Content,
			Members: addGroupParameter.Members,
		}
		newGroups = append(newGroups, newGroup)
	}
	response, err := g.collection.InsertMany(g.ctx, newGroups)
	if err != nil {
		panic(err)
	}
	return response
}

func (g *GroupsDB) UpdateGroupToDb(updateGroupParameter UpdateGroupParameter) *mongo.UpdateResult {
	getGroupResponse, getErr := g.GetGroupFromDb(updateGroupParameter.Filter)
	if getErr != nil {
		panic("Error: The original user do not exist.")
	}
	newGroup := Group{
		ID:      getGroupResponse.ID,
		Jackpot: updateGroupParameter.Jackpot,
		Title:   updateGroupParameter.Title,
		Content: updateGroupParameter.Content,
		Members: updateGroupParameter.Members,
	}

	updateUserResponse, err := g.collection.ReplaceOne(g.ctx, updateGroupParameter.Filter, newGroup)
	if err != nil {
		panic(err)
	}

	return updateUserResponse
}

func (g *GroupsDB) DeleteGroupFromDb(filter bson.D) *mongo.DeleteResult {
	response, err := g.collection.DeleteOne(g.ctx, filter)
	if err != nil {
		panic(err)
	}
	return response
}

func (g *GroupsDB) GetGroupsByUID(filter bson.D) (resultGroups []Group, err error) {
	// group list to return
	var groups []Group

	// find the groups whose memebers contain the specified uid
	cursor, _ := g.collection.Find(g.ctx, filter)

	if err = cursor.All(context.TODO(), &groups); err != nil {

		return nil, err
	}
	return groups, nil
}

func (g *GroupsDB) GetGroupMemberByGIDUID(gid, uid primitive.ObjectID) (member *Member, err error) {
	// First extract group contaning the member
	// then traverse group.members to find the member
	// needs optimization

	var group Group
	filter := bson.D{{"_id", gid}, {"members._id", uid}}
	err = g.collection.FindOne(g.ctx, filter).Decode(&group)
	if err != nil {
		return nil, err
	}
	for _, member := range group.Members {
		if member.UID == uid {
			return &member, nil
		}
	}
	return nil, errors.New("member not found")
}

func (g *GroupsDB) UpdateGroupMemberByGIDUID(newMember Member, gid, uid primitive.ObjectID) (*mongo.UpdateResult, error) {
	filter := bson.D{{"_id", gid}, {"members._id", uid}}
	result, err := g.collection.UpdateOne(g.ctx, filter, bson.D{
		{"$set", bson.D{{"members.$", newMember}}},
	})

	return result, err
}

func (g *GroupsDB) JoinGroupByGID(gid primitive.ObjectID, filter bson.D, member Member) error {
	// first test if the user already joined
	cnt, err := g.collection.CountDocuments(g.ctx, filter)
	if err != nil {
		return errors.New("joining group failed")
	}
	if cnt == 1 {
		return errors.New("already joined")
	}

	// have not joined yet, join
	update := bson.D{{"$push", bson.D{{"members", member}}}}
	res, err := g.collection.UpdateByID(g.ctx, gid, update)
	if err != nil {
		return err
	}
	if res.ModifiedCount == 1 {
		return nil
	} else {
		return errors.New("group not found")
	}
}

func (g *GroupsDB) QuitGroupByGID(gid, uid primitive.ObjectID, filter bson.D) error {
	cnt, err := g.collection.CountDocuments(g.ctx, filter)
	if err != nil {
		return errors.New("error while quitting group")
	}
	if cnt == 0 {
		return errors.New("bad request: user not joined")
	}
	// member are in the member list, so quit
	update := bson.D{{"$pull", bson.D{{"members", bson.D{{"_id", uid}}}}}}
	_, err = g.collection.UpdateByID(g.ctx, gid, update)
	if err != nil {
		return errors.New("error while updating database")
	}
	return nil
}
