package main

import (
	"flag"
	"os"

	"github.com/naveenkakumanu/go-crud-gin/models"
	router "github.com/naveenkakumanu/go-crud-gin/routers"
)

var port string

// passing app port from commandline argument with key eg: go run main.go port=8082
// Default Port is 8080 if we have not provided any port for go
// Now I can pass go run main.go port=<customport>
// Need to pass a string flag now accepting string
var ports = flag.String("port", ":8081", "Port to Execute the application")

// Reading Command Line Arguments with help of os and flag package

func main() {
	// To Read configuration from flag variables
	flag.Parse()

	// Reading command line argument using os package
	// if the values exists
	if len(os.Args) > 1 {
		port = os.Args[1]
	}

	// If the command line argument is missing reading from flag module
	if port == "" {
		port = *ports
	}

	// Initiating Router
	router := router.Router()

	config := models.DBConfig{}
	// Reading Yaml File
	config.ReadConfig()
	// Starting the Application
	router.Run(port)
}
