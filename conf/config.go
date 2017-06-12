package conf

import (
	"encoding/json"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// Neighbor config
type Neighbor struct {
	Address     string `json:"address"`
	Asn         uint32 `json:"asn"`
	Multihop    bool   `json:"multihop"`
	MultihopTTL uint8  `json:"multihopttl"`
}

// RawConfig config
type RawConfig map[string]interface{}

// Config the application's configuration
type Config struct {
	RouterID  string        `json:"routerid"`
	Asn       uint32        `json:"asn"`
	API       bool          `json:"api"`
	Neighbors []Neighbor    `json:"neighbors"`
	Logs      LoggingConfig `json:"logs"`
	Outputs   RawConfig     `json:"outputs"`
}

// LoadConfig loads the config from a file if specified, otherwise from the environment
func LoadConfig(cmd *cobra.Command) (*Config, error) {
	err := viper.BindPFlags(cmd.Flags())
	if err != nil {
		return nil, err
	}

	if configFile, _ := cmd.Flags().GetString("config"); configFile != "" {
		viper.SetConfigFile(configFile)
	} else {
		viper.SetConfigName("quasar")
		viper.AddConfigPath("./")
	}

	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}

	var config Config

	ko := viper.Unmarshal(&config)
	if ko != nil {
		return nil, ko
	}

	config.Outputs = viper.GetStringMap("outputs")

	return &config, nil
}

// ReflectConfig set conf from confraw
func ReflectConfig(confraw *interface{}, conf interface{}) (err error) {
	data, err := json.Marshal(confraw)
	if err != nil {
		return
	}

	if err = json.Unmarshal(data, conf); err != nil {
		return
	}

	return
}
