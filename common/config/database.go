package config

import "github.com/spf13/viper"

type DatabaseConfig struct {
	Host string
	Port int
	Username string
	Password string
	DBName string
}

func init() {
	viper.SetDefault("database.host","127.0.0.1")
	viper.SetDefault("database.port",5432)
	viper.SetDefault("database.username","postgres")
	viper.SetDefault("database.password","123456")
	viper.SetDefault("database.dbName","default")
}
