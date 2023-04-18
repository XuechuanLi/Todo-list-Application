package handlers

import (
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"net/http"
	"todo-app/primitives"
	"todo-app/stores"
)

type UsersHandler struct {
	db *stores.UsersDB
}

func NewUsersHandler(db *stores.UsersDB) *UsersHandler {
	return &UsersHandler{
		db,
	}
}

func (handler *UsersHandler) CreateUser(c *gin.Context) {
	var newUser primitives.AddUserParameter
	if err := c.ShouldBindJSON(&newUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	_, err := handler.db.AddUserToDb(newUser)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error while inserting a new user"})
		return
	}
	c.JSON(http.StatusOK, newUser)
}

func (handler *UsersHandler) UpdateUser(c *gin.Context) {
	uid, err := primitive.ObjectIDFromHex(c.Param("uid"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "invalid uid"})
		return
	}
	filter := bson.D{{"_id", uid}}

	var newUser primitives.UpdateUserParameter
	if err := c.ShouldBindJSON(&newUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid user"})
	}
	newUser.Filter = filter

	_, err = handler.db.UpdateUserToDb(newUser)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, gin.H{})
}

func (handler *UsersHandler) DeleteUser(c *gin.Context) {
	uid, err := primitive.ObjectIDFromHex(c.Param("uid"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "invalid uid"})
		return
	}

	filter := bson.D{{"_id", uid}}
	_, err = handler.db.DeleteUserFromDb(filter)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{})
}

func (handler *UsersHandler) ListUsers(c *gin.Context) {
	users := handler.db.BulkGetUsersFromDb(bson.D{})
	if len(users) == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "No users"})
		return
	}

	c.JSON(http.StatusOK, users)
}

func (handler *UsersHandler) ListUserByUID(c *gin.Context) {
	uid, err := primitive.ObjectIDFromHex(c.Param("uid"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "invalid uid"})
		return
	}
	filter := bson.D{{"_id", uid}}
	user, err := handler.db.GetUserFromDb(filter)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, user)
}
