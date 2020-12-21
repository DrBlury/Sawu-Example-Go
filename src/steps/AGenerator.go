package steps

import (
	"os"
	"reflect"
	"sawu-example-go/config"
	"sawu-example-go/entities"
	"strconv"
	"time"

	guuid "github.com/google/uuid"
	log "github.com/sirupsen/logrus"
)

type stubMapping map[string]interface{}

// StubStorage stores the function map
var StubStorage = stubMapping{}

// LoadStepFunctions loads the function context
func LoadStepFunctions() {
	StubStorage = map[string]interface{}{
		"StepOne":   StepOne,
		"StepTwo":   StepTwo,
		"StepThree": StepThree,
	}
}

// Call calls a specific function of the StubStorage
func Call(funcName string, params ...interface{}) (result interface{}, err error) {
	f := reflect.ValueOf(StubStorage[funcName])
	if len(params) != f.Type().NumIn() {
		log.Error("The number of params is out of index")
		return
	}
	in := make([]reflect.Value, len(params))
	for k, param := range params {
		in[k] = reflect.ValueOf(param)
	}
	var res []reflect.Value
	res = f.Call(in)
	result = res[0].Interface()
	if res[1].Interface() != nil {
		err = res[1].Interface().(error)
	}

	return
}

// GenerateNewNextStepEvent generates a new Event with UUID and timestamp.
func GenerateNewNextStepEvent(stepName string) entities.NextStepEvent {
	var nextStepEvent entities.NextStepEvent

	nextStepEvent.ID = guuid.New().String()
	nextStepEvent.ProcessInstanceID = guuid.New().String()
	//Set default broker ip if not set
	processName, isPresent := os.LookupEnv("sawu_process_name")
	if isPresent == false {
		processName = config.Defaults.Sawu.ProcessName
	}
	nextStepEvent.ProcessName = processName
	nextStepEvent.ProcessStep = stepName
	nextStepEvent.ProcessStepClass = stepName
	timestamp := strconv.FormatInt(time.Now().UnixNano(), 10)
	shortTimestamp := timestamp[0 : len(timestamp)-6]
	nextStepEvent.TimeStamp = shortTimestamp

	return nextStepEvent
}
