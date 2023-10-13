package main

import (

	"github.com/gin-gonic/gin"
)

 func main(){


 }

 type computerHardware struct{

	ID  string `json: "string"`
	Brand string `json: "string"`
	Type string `json: "string"`
	IsNew bool `json: "bool"`
 }