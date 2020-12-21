package steps

import (
	"sawu-example-go/entities"

	log "github.com/sirupsen/logrus"
)

// StepTwo is the second example step in the workflow.
func StepTwo(nextStepEvent entities.NextStepEvent) (entities.NextStepEvent, error) {
	newEvent := GenerateNewNextStepEvent("StepThree")
	newEvent.ComingFromID = nextStepEvent.ID
	newEvent.Data = newEvent.Data + " - this is a new Event from Step two"
	log.Info("I'm in step two.")
	return newEvent, nil
}
