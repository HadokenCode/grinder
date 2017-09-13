package config

import (
	"fmt"

	"github.com/spf13/viper"
)

// Load will read config (config.yaml) file into viper struct
func Load(dirs ...string) *viper.Viper {
	v := viper.New()
	v.SetConfigName("config")
	v.AddConfigPath(".")

	// add paths to search
	for _, dir := range dirs {
		v.AddConfigPath(dir)
	}

	err := v.ReadInConfig() // Find and read the config file
	if err != nil {         // Handle errors reading the config file
		panic(fmt.Errorf("fatal error config file: %s", err))
	}
	return v
}
