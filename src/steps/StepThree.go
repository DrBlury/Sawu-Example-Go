package steps

import (
	"sawu-example-go/entities"

	log "github.com/sirupsen/logrus"
)

// StepThree is the third example step in the workflow.
func StepThree(nextStepEvent entities.NextStepEvent) (entities.NextStepEvent, error) {
	newEvent := GenerateNewNextStepEvent("SAWUEND")
	newEvent.ComingFromID = nextStepEvent.ID
	newEvent.Data = newEvent.Data + " - this is a new Event from Step three"
	log.Info("I'm in step three.")
	return newEvent, nil
}
