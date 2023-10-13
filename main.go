package main

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

// define JSON strucutre
type computerHardware struct {
	ID    string `json: "string"`
	Brand string `json: "string"`
	Type  string `json: "string"`
	Model string `json: "string"`
	IsNew bool   `json: "bool"`
}

// Array to store data
var computers = []computerHardware{
	{
		ID:    "21",
		Brand: "MSI",
		Type:  "GPU",
		Model: "RTX 4080ti",
		IsNew: false,
	},
}

// Starting the server and define routes
func main() {

	// Creating routes for all functions
	router := gin.Default()
	router.GET("/getAll", GetAll)
	router.GET("/get/:id", GetbyID)
	router.POST("/key/set", addToMemory)

	//Start server on port 9000
	router.Run(":9000")
}

//Define all functionalities needed in given routes

// Create a GetAll function to fetch all entires to the DB
func GetAll(context *gin.Context) {

	context.IndentedJSON(http.StatusOK, computers)
}

// Get element by particular ID
func GetbyID(context *gin.Context) {
	id:= context.Param("id")
	computer, err := getComputerbyID(id)

	if err != nil {
		return
	}

	context.IndentedJSON(http.StatusOK, computer)
}
	// function to scan array for particular ID
func getComputerbyID(id string)(*computerHardware, error){
	for i, comp := range computers{
		if comp.ID == id{
			return &computers[i], nil
		}
	}
	return nil, errors.New("Hardware with given ID not found")
}

// Handle adding new entries to array. 
// POST needs to be in form of an array. Please make sure the JSON format is wrapped in an array [{}] otherwise 
// it will count as a invalid body.
func addToMemory(context *gin.Context){

	//var newHardware computerHardware 
	var newHardwareArr[] computerHardware
	// Validate JSON body with struct for handling error. If body is invalid, throw an error
	if err:= context.BindJSON(&newHardwareArr); err != nil {
		return
	}

	// Add the new data to existing array
	computers = append(computers, newHardwareArr...)
	context.IndentedJSON(http.StatusCreated, newHardwareArr)

} 

