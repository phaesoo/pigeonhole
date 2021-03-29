package configs

import (
	"fmt"

	"github.com/spf13/viper"
)

// AppConfig is config struct for app
type AppConfig struct {
	Name    string
	Host    string
	Port    int
	Profile bool
	Metrics bool
}

func appConfig(v *viper.Viper) AppConfig {
	return AppConfig{
		Name:    v.GetString("pigeonhole.name"),
		Host:    v.GetString("pigeonhole.host"),
		Port:    v.GetInt("pigeonhole.port"),
		Profile: v.GetBool("pigeonhole.profile"),
		Metrics: v.GetBool("pigeonhole.metrics"),
	}
}

func (c *AppConfig) Address() string {
	return fmt.Sprintf("%s:%d", c.Host, c.Port)
}
