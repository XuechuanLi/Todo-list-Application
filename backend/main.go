// Todo-app API
//
// This is todo-app API.
//
//		Schemes: http
//	 Host: localhost:8080
//		BasePath: /
//		Version: 0.0.1
//		Contact: Zhongtao Miao <zhongtao.miao@gmail.com> http://127.0.0.1:8080
//
//		Consumes:
//		- application/json
//
//		Produces:
//		- application/json
//
// swagger:meta
package main

import (
	"context"
	"github.com/gin-gonic/gin"
	"todo-app/handlers"
	"todo-app/stores"
	"todo-app/utils"
)

var (
	groupsHandler *handlers.GroupsHandler
	usersHandler  *handlers.UsersHandler
)

func init() {
	groupsHandler = handlers.NewGroupsHandler(
		stores.NewGroupsDB(
			context.TODO(),
			utils.GetCollection("MONGO_DATABASE", "MONGO_GROUP_COLLECTION"),
		),
	)
	usersHandler = handlers.NewUsersHandler(
		stores.NewUsersDB(
			context.TODO(),
			utils.GetCollection("MONGO_DATABASE", "MONGO_USER_COLLECTION"),
		),
	)
}
func main() {
	router := gin.Default()

	router.GET("/groups", groupsHandler.ListAllGroups)
	router.POST("/groups", groupsHandler.CreateGroup)
	router.DELETE("/groups/:gid", groupsHandler.DeleteGroupByGID)
	router.PUT("/groups/:gid", groupsHandler.UpdateGroupByGID)

	router.GET("/groups/:uid", groupsHandler.GetGroupsByUID)
	router.PUT("/groups/join", groupsHandler.JoinGroupByGID)
	router.PUT("/groups/quit", groupsHandler.QuitGroupByGID)

	router.GET("/member/:gid/:uid", groupsHandler.GetGroupMemberByGIDUID)
	router.PUT("/member/:gid/:uid", groupsHandler.UpdateGroupMemberByGIDUID)

	router.GET("/users", usersHandler.ListUsers)
	router.GET("/users/:uid", usersHandler.ListUserByUID)
	router.POST("/users", usersHandler.CreateUser)
	router.PUT("/users/:uid", usersHandler.UpdateUser)
	router.DELETE("/users/:uid", usersHandler.DeleteUser)

	router.Run("localhost:8080")
}
