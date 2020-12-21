package main

import (
	"os"
	"sawu-example-go/config"
	"sawu-example-go/kafka"
	"sawu-example-go/steps"

	"github.com/gofiber/fiber/v2"
)

var app *fiber.App

func main() {
	var fiberConfig = fiber.Config{}
	fiberConfig.DisableStartupMessage = false
	app = fiber.New(fiberConfig)
	config.GetDefaults()
	addRestController()
	steps.LoadStepFunctions()
	go kafka.DoKafkaConsumerStuff()
	kafka.CreateProducer()

	//Set default port if not set
	port, isPresent := os.LookupEnv("fiber_port")
	if isPresent == false {
		port = config.Defaults.Port
	}
	//kafka.CreateProducer()
	app.Listen(":" + port)
}

// AddRestController creates the controller for all sorts of things
func addRestController() {

}
