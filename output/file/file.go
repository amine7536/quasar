package outputfile

import (
	"encoding/json"
	"os"

	"github.com/Sirupsen/logrus"
	"github.com/amine7536/quasar/conf"
	"github.com/amine7536/quasar/event"
)

// Name output name
const Name = "file"

// Output main type
type Output struct {
	File   string `json:"file"`
	Pretty bool   `json:"pretty"`
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

// Send send output
func (t *Output) Send(event event.Event) error {
	f, err := os.OpenFile(t.File, os.O_CREATE|os.O_RDWR|os.O_APPEND, 0660)
	if err != nil {
		return err
	}

	var output []byte
	if t.Pretty == true {
		output, err = json.MarshalIndent(event, "    ", "  ")
	} else {
		output, err = json.Marshal(event)
	}

	if err != nil {
		return err
	}

	f.Write(output)

	return nil
}
