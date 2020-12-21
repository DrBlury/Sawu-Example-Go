package main

import (
	"fmt"
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
	app.Get("/start/:data", func(c *fiber.Ctx) error {
		// Generate new Event
		nextStepEvent := steps.GenerateNewNextStepEventForNewProcess("StepOne")
		nextStepEvent.Data = c.Params("data")
		// Send Event to Kafka
		kafka.SendNextStepEvent(nextStepEvent)

		return c.Status(202).JSON(&fiber.Map{
			"message": fmt.Sprintf("Accepted. Started processing with params: %s.", c.Params("data")),
		})
	})
}
