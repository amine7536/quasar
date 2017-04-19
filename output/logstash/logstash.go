package outputlogstash

import (
	"encoding/json"
	"fmt"

	"github.com/Sirupsen/logrus"
	"github.com/amine7536/quasar/conf"
	"github.com/amine7536/quasar/event"
)

// Name output name
const Name = "logstash"

// OutputConfig main type
type Output struct {
	Host string `json:"host"`
	Port string `json:"port"`
}

// DefaultOutputConfig return default config struct
func DefaultOutputConfig() Output {
	return Output{}
}

// InitOutput initialize an output
func InitOutput(configRaw *interface{}, logger *logrus.Entry) (conf.OutputHandler, error) {
	config := DefaultOutputConfig()
	err := conf.ReflectConfig(configRaw, &config)
	if err != nil {
		return &config, err
	}

	return &config, err
}

// Output send output
func (t *Output) Send(event event.Event) error {
	output, err := json.MarshalIndent(event, "    ", "  ")

	if err != nil {
		fmt.Println(err)
		return err
	}

	fmt.Println(string(output))
	return nil
}
