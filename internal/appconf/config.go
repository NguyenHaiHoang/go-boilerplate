package appconf

import (
	"apus-sample/common/config"
)

var root AppConfig

type AppConfig struct {
	Database    config.DatabaseConfig
	Transporter TransporterConfig
}

func Database() config.DatabaseConfig {
	return root.Database
}

func Transporter() TransporterConfig {
	return root.Transporter
}

func LoadConfig() error {
	return config.LoadConfig(&root)
}
