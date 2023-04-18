package stores

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
	. "todo-app/models"
	. "todo-app/primitives"
)

//var (
//	userCollection = utils.GetCollection("MONGO_DATABASE", "MONGO_USER_COLLECTION")
//)

type UsersDB struct {
	collection *mongo.Collection
	ctx        context.Context
}

func NewUsersDB(ctx context.Context, collection *mongo.Collection) *UsersDB {
	return &UsersDB{
		collection: collection,
		ctx:        ctx,
	}
}

func (u *UsersDB) GetUserFromDb(filter bson.D) (resultUser *User, err error) {
	err = u.collection.FindOne(u.ctx, filter).Decode(&resultUser)

	if err != nil {
		if err == mongo.ErrNoDocuments {
			return
		}
		panic(err)
	}
	return resultUser, nil
}

func (u *UsersDB) BulkGetUsersFromDb(filter bson.D) (resultUsers []User) {
	bulkGetResponse, err := u.collection.Find(u.ctx, filter)
	if err != nil {
		log.Fatal(err)
	}

	defer func(bulkGetResponse *mongo.Cursor, ctx context.Context) {
		err := bulkGetResponse.Close(ctx)
		if err != nil {

		}
	}(bulkGetResponse, u.ctx)

	for bulkGetResponse.Next(u.ctx) {

		var result User
		err := bulkGetResponse.Decode(&result)
		if err != nil {
			log.Fatal(err)
		}
		resultUsers = append(resultUsers, result)
	}
	if err := bulkGetResponse.Err(); err != nil {
		log.Fatal(err)
	}
	return resultUsers
}

func (u *UsersDB) AddUserToDb(newUserParameters AddUserParameter) (*mongo.InsertOneResult, error) {
	newUser := User{
		ID:          primitive.NewObjectID(),
		Username:    newUserParameters.Username,
		Email:       newUserParameters.Email,
		Password:    newUserParameters.Password,
		TotalPoints: newUserParameters.TotalPoints,
		CreatedAt:   newUserParameters.CreatedAt,
	}
	response, err := u.collection.InsertOne(u.ctx, newUser)
	if err != nil {
		return nil, err
	}
	return response, nil
}

func (u *UsersDB) BulkAddUsersToDb(bulkAddUsersParameters BulkAddUsersParameters) *mongo.InsertManyResult {
	var newUsers []interface{}
	for index := 0; index < len(bulkAddUsersParameters.Parameters); index++ {
		addUserParameter := bulkAddUsersParameters.Parameters[index]
		newUser := User{
			ID:          primitive.NewObjectID(),
			Username:    addUserParameter.Username,
			Email:       addUserParameter.Email,
			Password:    addUserParameter.Password,
			TotalPoints: addUserParameter.TotalPoints,
			CreatedAt:   addUserParameter.CreatedAt,
		}
		newUsers = append(newUsers, newUser)
	}
	response, err := u.collection.InsertMany(u.ctx, newUsers)
	if err != nil {
		panic(err)
	}
	return response
}

func (u *UsersDB) UpdateUserToDb(updateUserParameter UpdateUserParameter) (*mongo.UpdateResult, error) {
	getUserResponse, getErr := u.GetUserFromDb(updateUserParameter.Filter)
	if getErr != nil {
		panic("Error: The original user do not exist.")
	}
	newUser := User{
		ID:          getUserResponse.ID,
		Username:    updateUserParameter.Username,
		Email:       updateUserParameter.Email,
		Password:    updateUserParameter.Password,
		TotalPoints: updateUserParameter.TotalPoints,
		CreatedAt:   updateUserParameter.CreatedAt,
	}

	updateUserResponse, err := u.collection.ReplaceOne(u.ctx, updateUserParameter.Filter, newUser)
	if err != nil {
		return nil, err
	}

	return updateUserResponse, nil
}

func (u *UsersDB) DeleteUserFromDb(filter bson.D) (*mongo.DeleteResult, error) {
	response, err := u.collection.DeleteOne(u.ctx, filter)
	if err != nil {
		return nil, err
	}
	return response, nil
}
