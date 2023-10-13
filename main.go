package main

import (

	"github.com/gin-gonic/gin"
	"net/http"
)
//define JSON strucutre
type computerHardware struct{

	ID  string `json: "string"`
	Brand string `json: "string"`
	Type string `json: "string"`
	Model string `json: "string"`
	IsNew bool `json: "bool"`
 }
 
// Array to store data
 var computers = []computerHardware{
	{
		ID: "ckw26",
		Brand: "MSI",
		Type: "GPU",
		Model: "RTX 4080ti",
		IsNew: false,
	},  
}

//Starting the server and define routes
 func main(){

	// Creating routes for all functions
	router := gin.Default()
	router.GET("/get", GetAll)
	
	//Start server on port 9000
	router.Run(":9000")
 }
//Define all functionalities needed in given routes

// Create a GetAll function to fetch all entires to the DB
func GetAll(context *gin.Context)(){

	context.IndentedJSON(http.StatusOK, computers)
}