package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/yurinohana/ApiServer/db"
	"github.com/yurinohana/ApiServer/models"
)

func Index(c *gin.Context) {
	//user := models.User{ID:1,Name:"yuria"}
	//users := []models.User{user}
	//c.JSON(200, users)

	db := db.DynamodbConnect()
	table := db.Table("Users")

	var users []models.User
	err := table.Scan().All(&users)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(users)
}

func Show(c *gin.Context) {
	//n := c.Param("id")
	//id, _ := strconv.Atoi(n)
	//user := models.User{
	//	ID: id,
	//}
	//c.JSON(200, user)

	db := db.DynamodbConnect()
	table := db.Table("Users")

	var user models.User
	err := table.Get("id", 1).One(&user)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(user)
}

func Create(c *gin.Context) {
	newUser := models.User{ID:3,Name:"yuria"}
	//id, _ := strconv.Atoi(c.PostForm("id"))
	//newUser.ID = id
	//newUser.Name = c.PostForm("Name")
	//c.JSON(200, newUser)

	db := db.DynamodbConnect()
	table := db.Table("Users")

	err := table.Put(newUser)
	if err != nil {
		fmt.Println(err)
	}
}

func Delete(c *gin.Context) {
	//n := c.Param("id")
	//id, _ := strconv.Atoi(n)
	//user := models.User{ID:id}
	//c.JSON(200, "")

	db := db.DynamodbConnect()
	table := db.Table("Users")

	err := table.Delete("id", 2)
	if err != nil {
		fmt.Println(err)
	}
}