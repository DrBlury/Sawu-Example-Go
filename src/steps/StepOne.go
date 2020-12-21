package steps

import (
	"sawu-example-go/entities"

	log "github.com/sirupsen/logrus"
)

// StepOne is the first example step in the workflow.
func StepOne(nextStepEvent entities.NextStepEvent) (entities.NextStepEvent, error) {
	newEvent := GenerateNewNextStepEvent("StepTwo")
	newEvent.ComingFromID = nextStepEvent.ID
	newEvent.Data = "Banana..."
	log.Info("I'm in step one.")
	return newEvent, nil
}
