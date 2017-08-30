package outputstdout

import (
	"encoding/json"
	"fmt"

	"github.com/amine7536/quasar/conf"
	"github.com/amine7536/quasar/event"
	"github.com/sirupsen/logrus"
)

// Name output name
const Name = "stdout"

// OutputConfig main type
type Output struct {
	Mode string `json:"mode"`
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
	if t.Mode == "json" {
		output, err := json.MarshalIndent(event, "    ", "  ")

		if err != nil {
			fmt.Println(err)
			return err
		}

		fmt.Println(string(output))
	} else {
		output, err := json.Marshal(event)

		if err != nil {
			fmt.Println(err)
			return err
		}

		fmt.Println(output)
	}
	return nil
}
