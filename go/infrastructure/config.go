package infrastructure

import (
	"fmt"
	"os"
	"strings"

	"github.com/spf13/viper"
	"go.uber.org/zap"
)

const (
	prefix = "LIAISON"
)

type Config struct {
	MongoSettings struct {
		ConnectionString string `json:"ConnectionString"`
		Database         string `json:"Database"`
	} `json:"MongoSettings"`
	Host struct {
		Port int `json:"Port"`
	} `json:"Host"`
}

// Populate reads the configuration from the environment and config file
func (conf *Config) Populate(logger *zap.Logger) {
	logger.Info("Reading configuration")
	viper.SetEnvPrefix(prefix)
	viper.AutomaticEnv()

	// inspired by https://github.com/spf13/viper/issues/761#issuecomment-1578931559
	for _, e := range os.Environ() {
		split := strings.Split(e, "=")
		k := split[0]
		if strings.HasPrefix(k, prefix) {
			name := strings.Join(strings.Split(k, "_")[1:], ".")
			// Explicit Set has the highest priority
			viper.Set(name, split[1])
		}
	}

	viper.SetConfigName("config")
	viper.AddConfigPath(".")
	err := viper.ReadInConfig()
	if err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			logger.Warn("Config file not found, using environment variables")
		} else {
			panic(fmt.Errorf("config file error %w", err))
		}
	}

	err = viper.Unmarshal(&conf)
	if err != nil {
		panic(fmt.Errorf("config unmarshalling error %w", err))
	}
}
