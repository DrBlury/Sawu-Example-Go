package steps

import (
	"sawu-example-go/entities"

	log "github.com/sirupsen/logrus"
)

// StepOne is the first example step in the workflow.
func StepOne(oldStepEvent entities.NextStepEvent) (entities.NextStepEvent, error) {
	newEvent := GenerateNewNextStepEvent(oldStepEvent, "StepTwo")
	newEvent.Data = "Banana..."
	log.Info("I'm in step one.")
	return newEvent, nil
}
