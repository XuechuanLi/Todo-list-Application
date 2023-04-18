package handlers

import (
	"context"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
	"net/http"
	"todo-app/models"
	"todo-app/primitives"
	"todo-app/stores"
	"todo-app/utils"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
)

type GroupsHandler struct {
	db *stores.GroupsDB
}

func NewGroupsHandler(db *stores.GroupsDB) *GroupsHandler {
	return &GroupsHandler{
		db,
	}
}

func (handler *GroupsHandler) ListAllGroups(c *gin.Context) {
	//c.IndentedJSON(http.StatusOK, groups)
	//log.Printf("Request to MongoDB")
	//cur, err := handler.collection.Find(handler.ctx, bson.M{})
	//if err != nil {
	//	c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	//	return
	//}
	//defer cur.Close(handler.ctx)
	//
	//groups := make([]models.Group, 0)
	//for cur.Next(handler.ctx) {
	//	var group models.Group
	//	cur.Decode(&group)
	//	groups = append(groups, group)
	//}
	//
	//c.JSON(http.StatusOK, groups)
	groups, err := handler.db.BulkGetGroupsFromDb(bson.D{})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	if len(groups) == 0 {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "No groups"})
		return
	}
	c.JSON(http.StatusOK, groups)
}

func (handler *GroupsHandler) CreateGroup(c *gin.Context) {
	//var newGroup models.Group
	//
	//// bind the received group json to newGroup
	//if err := c.BindJSON(&newGroup); err != nil {
	//	c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	//	return
	//}
	//
	//// add the newGroup
	//groups = append(groups, newGroup)
	//// created successfully
	//c.IndentedJSON(http.StatusCreated, newGroup)

	var newGroup primitives.AddGroupParameter
	if err := c.ShouldBindJSON(&newGroup); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	//group.ID = primitive.NewObjectID()
	//// gurantee that members is of type array
	//// query in db causes error when it's of type null
	//group.Members = []models.Member{}

	_, err := handler.db.AddGroupToDb(newGroup)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error while inserting a new group"})
		return
	}
	c.JSON(http.StatusOK, newGroup)
}

func (handler *GroupsHandler) DeleteGroupByGID(c *gin.Context) {
	//gidStr := c.Param("gid")
	//gid, err := strconv.Atoi(gidStr)
	//if err != nil {
	//	c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "invalid gid"})
	//}
	//
	//for indexGroup, group := range groups {
	//	if group.GID == gid {
	//		groups = append(groups[:indexGroup], groups[indexGroup+1:]...)
	//		c.IndentedJSON(http.StatusOK, gin.H{"message": "delete group successfully"})
	//		return
	//	}
	//}
	//c.IndentedJSON(http.StatusNotFound, gin.H{"error": "group not found"})

	gidStr := c.Param("gid")
	gid, _ := primitive.ObjectIDFromHex(gidStr)
	filter := bson.D{{"_id", gid}}
	result := handler.db.DeleteGroupFromDb(filter)

	if result.DeletedCount == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "group not found!",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "deleted successfully"})
}

func (handler *GroupsHandler) UpdateGroupByGID(c *gin.Context) {
	//gidStr := c.Param("gid")
	//gid, err := strconv.Atoi(gidStr)
	//if err != nil {
	//	c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "invalid gid"})
	//}
	//
	//var newGroup models.Group
	//if err := c.BindJSON(&newGroup); err != nil {
	//	return
	//}
	//
	//for indexGroup, group := range groups {
	//	if group.GID == gid {
	//		groups[indexGroup] = newGroup
	//		c.IndentedJSON(http.StatusOK, newGroup)
	//		return
	//	}
	//}
	//c.IndentedJSON(http.StatusNotFound, gin.H{"error": "group not found"})

	gidStr := c.Param("gid")
	gid, _ := primitive.ObjectIDFromHex(gidStr)
	filter := bson.D{{"_id", gid}}

	var group primitives.UpdateGroupParameter
	if err := c.ShouldBindJSON(&group); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	group.Filter = filter

	result := handler.db.UpdateGroupToDb(group)

	if result == nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error while updating a group"})
		return
	}
	c.JSON(http.StatusOK, group)
}

func (handler *GroupsHandler) GetGroupsByUID(c *gin.Context) {
	uid, err := primitive.ObjectIDFromHex(c.Param("uid"))
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{
			"message": "invalid uid"})
		return
	}

	filter := bson.D{{"members._id", uid}}
	groups, err := handler.db.GetGroupsByUID(filter)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error while finding groups by uid"})
		return
	}
	c.IndentedJSON(http.StatusOK, groups)
}

func (handler *GroupsHandler) GetGroupMemberByGIDUID(c *gin.Context) {
	gid, err := primitive.ObjectIDFromHex(c.Param("gid"))
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{
			"message": "invalid gid"})
		return
	}

	uid, err := primitive.ObjectIDFromHex(c.Param("uid"))
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{
			"message": "invalid uid"})
		return
	}

	log.Printf("Get Group Member By GID = %v, UID = %v\n", gid, uid)

	member, err := handler.db.GetGroupMemberByGIDUID(gid, uid)
	if err == mongo.ErrNoDocuments {
		c.JSON(http.StatusNotFound, gin.H{"error": "group not found or member not in group"})
	} else if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error while finding group by uid and gid"})
	}
	c.JSON(http.StatusOK, member)
}

func (handler *GroupsHandler) UpdateGroupMemberByGIDUID(c *gin.Context) {
	gid, err := primitive.ObjectIDFromHex(c.Param("gid"))
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{
			"message": "invalid gid"})
		return
	}

	uid, err := primitive.ObjectIDFromHex(c.Param("uid"))
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{
			"message": "invalid uid"})
		return
	}

	log.Printf("update Group Member By GID = %v, UID = %v\n", gid, uid)

	var newMember models.Member
	if err := c.ShouldBindJSON(&newMember); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "invalid post body " + err.Error()})
		return
	}

	result, err := handler.db.UpdateGroupMemberByGIDUID(newMember, gid, uid)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error while finding group by uid and gid " + err.Error()})
		return
	}
	if result.ModifiedCount == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "Group or member not found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{})
}

func (handler *GroupsHandler) JoinGroupByGID(c *gin.Context) {
	gid, err := primitive.ObjectIDFromHex(c.Query("gid"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "invalid gid"})
		return
	}

	var member models.Member
	if err := c.ShouldBindJSON(&member); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error()})
		return
	}

	// We need change the user's total_points as well
	// If the user doesn't have enough points, he/she can't join the group

	uid := member.UID
	userFilter := bson.D{{"_id", uid}}
	usersHandler := NewUsersHandler(
		stores.NewUsersDB(
			context.TODO(),
			utils.GetCollection("MONGO_DATABASE", "MONGO_USER_COLLECTION"),
		),
	)
	user, err := usersHandler.db.GetUserFromDb(userFilter)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	if user.TotalPoints > member.Points {
		user.TotalPoints -= member.Points
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"error": "the user does not have enough points"})
	}
	var updateUser primitives.UpdateUserParameter
	updateUser.Username = user.Username
	updateUser.Email = user.Email
	updateUser.Password = user.Password
	updateUser.CreatedAt = user.CreatedAt
	updateUser.TotalPoints = user.TotalPoints
	updateUser.Filter = userFilter

	_, err = usersHandler.db.UpdateUserToDb(updateUser)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	filter := bson.D{{"_id", gid}, {"members._id", member.UID}}
	err = handler.db.JoinGroupByGID(gid, filter, member)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// When we arrive here, the new member has joined.
	// However, we need to change the current group's jackpot as well
	groupFilter := bson.D{{"_id", gid}}

	// get the points of new joint user
	point := member.Points
	// get current group
	currentGroup, err := handler.db.GetGroupFromDb(groupFilter)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	// update current group's jackpot
	currentGroup.Jackpot += point

	// update current group
	var group primitives.UpdateGroupParameter
	group.Title = currentGroup.Title
	group.Jackpot = currentGroup.Jackpot
	group.Content = currentGroup.Content
	group.Members = currentGroup.Members
	group.Filter = groupFilter

	result := handler.db.UpdateGroupToDb(group)

	if result == nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error while updating a group"})
		return
	}
	c.JSON(http.StatusOK, gin.H{})
}

func (handler *GroupsHandler) QuitGroupByGID(c *gin.Context) {
	gid, err := primitive.ObjectIDFromHex(c.Query("gid"))
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{
			"message": "invalid gid"})
		return
	}

	uid, err := primitive.ObjectIDFromHex(c.Query("uid"))
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{
			"message": "invalid uid"})
		return
	}

	// Before the user quit successfully, we need to reduce the current group's jackpot respectively
	// first get member
	member, err := handler.db.GetGroupMemberByGIDUID(gid, uid)
	if err == mongo.ErrNoDocuments {
		c.JSON(http.StatusNotFound, gin.H{"error": "group not found or member not in group"})
	} else if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error while finding group by uid and gid"})
	}
	c.JSON(http.StatusOK, member)
	// get current group
	groupFilter := bson.D{{"_id", gid}}
	// get the points of the quit user
	point := member.Points
	// get current group
	currentGroup, err := handler.db.GetGroupFromDb(groupFilter)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	// update current group's jackpot
	currentGroup.Jackpot -= point
	// update current group
	var group primitives.UpdateGroupParameter
	group.Title = currentGroup.Title
	group.Jackpot = currentGroup.Jackpot
	group.Content = currentGroup.Content
	group.Members = currentGroup.Members
	group.Filter = groupFilter

	result := handler.db.UpdateGroupToDb(group)
	if result == nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error while updating a group"})
		return
	}

	// quit code
	filter := bson.D{{"_id", gid}, {"members._id", uid}}
	err = handler.db.QuitGroupByGID(gid, uid, filter)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	// Todo: When the user quits, how to deal with his/her points in this group?

	c.JSON(http.StatusOK, gin.H{})
}
