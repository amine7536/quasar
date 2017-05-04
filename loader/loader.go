package loader

import (
	"fmt"

	"github.com/Sirupsen/logrus"
	"github.com/amine7536/quasar/conf"
	"github.com/amine7536/quasar/output/file"
	"github.com/amine7536/quasar/output/logstash"
	"github.com/amine7536/quasar/output/stdout"
)

// LoadOutputs load outputs
func LoadOutputs(outputConfig conf.RawConfig, logger *logrus.Entry) error {
	// Register Outputs
	if outputConfig != nil {
		for k, v := range outputConfig {

			switch k {
			case "stdout":
				output, err := outputstdout.InitOutput(&v, logger)
				conf.RegisterOutput(k, output)
				if err != nil {
					return fmt.Errorf("Faild to load output %s", k)
				}

			case "logstash":
				output, err := outputlogstash.InitOutput(&v, logger)
				conf.RegisterOutput(k, output)
				if err != nil {
					return fmt.Errorf("Faild to load output %s", k)
				}

			case "file":
				output, err := outputfile.InitOutput(&v, logger)
				conf.RegisterOutput(k, output)
				if err != nil {
					return fmt.Errorf("Faild to load output %s", k)
				}

			}
		}
	}

	return nil
}
