package conf

import (
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// Neighbor config
type Neighbor struct {
	Address string `json:"address"`
	Asn     uint32 `json:"asn"`
}

// Outputs config
type Outputs struct {
	Mode string `json:"mode"`
	Host string `json:"host"`
	Port string `json:"port"`
}

// Config the application's configuration
type Config struct {
	RouterID  string        `json:"routerid"`
	Asn       uint32        `json:"asn"`
	API       bool          `json:"api"`
	Neighbors []Neighbor    `json:"neighbors"`
	Logs      LoggingConfig `json:"logs"`
	Outputs   []Outputs     `json:"outputs"`
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

	return &config, nil
}
