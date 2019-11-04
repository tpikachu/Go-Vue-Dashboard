package config

import (
	"fmt"

	"github.com/spf13/viper"
)

// OktaConfig Config for okta
type OktaConfig struct {
	ClientID string `mapstructure:"client_id"`
	Issuer   string `mapstructure:"issuer"`
}

// VaultServerConfig Config for vault servers
type VaultServerConfig struct {
	Name  string `mapstructure:"name"`
	Addr  string `mapstructure:"addr"`
	Token string `mapstructure:"token"`
}

// Config Create private data struct to hold config options.
type Config struct {
	Okta        OktaConfig          `mapstructure:"okta"`
	VaultServer []VaultServerConfig `mapstructure:"vault_servers"`
}

// Create a new config instance.
var (
	Conf *Config
)

// Read the config file from the current directory and marshal
// into the conf config struct.
func getConf() *Config {
	viper.AddConfigPath(".")
	viper.SetConfigName("config")
	err := viper.ReadInConfig()

	if err != nil {
		fmt.Printf("%v", err)
	}

	conf := &Config{}
	err = viper.Unmarshal(conf)
	if err != nil {
		fmt.Printf("unable to decode into config struct, %v", err)
	}

	return conf
}

// Initialization routine.
func init() {
	// Retrieve config options.
	Conf = getConf()
}
