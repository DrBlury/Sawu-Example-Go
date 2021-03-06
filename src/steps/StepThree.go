package steps

import (
	"sawu-example-go/entities"

	log "github.com/sirupsen/logrus"
)

// StepThree is the third example step in the workflow.
func StepThree(oldStepEvent entities.NextStepEvent) (entities.NextStepEvent, error) {
	newEvent := GenerateNewNextStepEvent(oldStepEvent, "SAWUEND")
	if len(oldStepEvent.Data) != 0 {
		newEvent.Data = oldStepEvent.Data + " - this is a new Event from Step three"
	}
	log.Info("I'm in step three.")
	return newEvent, nil
}
