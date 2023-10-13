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

func getComputerbyID(id string)(*computerHardware, error){
	for i, comp := range computers{
		if comp.ID == id{
			return &computers[i], nil
		}
	}
	return nil, errors.New("Hardware with given ID not found")
}
