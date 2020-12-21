package config

import (
	"io/ioutil"

	log "github.com/sirupsen/logrus"
	"gopkg.in/yaml.v2"
)

// Conf is the struct that holds configuration data
type Conf struct {
	Port string `yaml:"port"`
	Sawu struct {
		Separator   string `yaml:"separator"`
		ProcessName string `yaml:"processname"`
	}
	Kafka struct {
		Broker struct {
			IPAddress string `yaml:"ipaddress"`
		}
		Consumer struct {
			ConsumerGroup string   `yaml:"consumergroup"`
			Topics        []string `yaml:"topics"`
		}
	}
}

// Defaults struct for project wide usage
var Defaults Conf

// GetDefaults loads the default configuration
func GetDefaults() {

	yamlFile, err := ioutil.ReadFile("../res/defaults.yml")
	if err != nil {
		log.Printf("yamlFile.Get err   #%v ", err)
	}
	err = yaml.Unmarshal(yamlFile, &Defaults)
	if err != nil {
		log.Fatalf("Unmarshal: %v", err)
	}
}
