package config

import "github.com/spf13/viper"

type RestServerConfig struct {
	Host string
	Port int
}

func init() {
	viper.SetDefault("transporter.rest.host","0.0.0.0")
	viper.SetDefault("transporter.rest.port",80)
}
