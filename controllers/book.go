package controllers

import (
	"github.com/gin-gonic/gin"
	"main.go/db"
	"main.go/entity"
)

var Db = db.ConnectDB()

// Find All Data
func GetAll(context *gin.Context) {
	var data []entity.Add_e
	Db.Table("purpose_of_cost").Find(&data)
	//Db.Save(&data)
	context.IndentedJSON(200, data)

}

//Create Note your costing purpose
func Creates(context *gin.Context) {

	u := entity.Add_e{}
	err := context.ShouldBindJSON(&u)
	//fmt.Println(u)
	if err != nil {
		context.IndentedJSON(200, gin.H{"Message": "Error"})
	} else {
		context.IndentedJSON(200, gin.H{"Message": "Successfully Created your Work file"})
	}
	Db.Table("purpose_of_cost").Create(&u)
	context.IndentedJSON(200, u)

}

//Delete Your note which is done
func Deletes(context *gin.Context) {
	var err error
	n := entity.Naam{}
	context.ShouldBindJSON(&n)
	de := `delete from "purpose_of_cost" where id=$1`
	Db.Exec(de, n.Id)
	if err != nil {
		context.IndentedJSON(200, gin.H{"Message": "Error"})
	} else {
		context.IndentedJSON(200, gin.H{"Message": "Successfully Created your Work file"})
	}
	context.IndentedJSON(200, n)

}
