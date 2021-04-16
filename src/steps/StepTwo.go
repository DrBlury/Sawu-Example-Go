package steps

import (
	"sawu-example-go/entities"

	log "github.com/sirupsen/logrus"
)

// StepTwo is the second example step in the workflow.
func StepTwo(oldStepEvent entities.NextStepEvent) (entities.NextStepEvent, error) {
	newEvent := GenerateNewNextStepEvent(oldStepEvent, "StepThree")
	if len(oldStepEvent.Data) != 0 {
		newEvent.Data = oldStepEvent.Data + " - this is a new Event from Step two"
	}
	log.Info("I'm in step two.")
	return newEvent, nil
}
