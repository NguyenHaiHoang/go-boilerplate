package config

import (
	"apus-sample/common/utils"
	"fmt"
	"github.com/spf13/viper"
)

func LoadConfig(confObj interface{})error{
	configFileName := fmt.Sprintf("configfiles/%s.yaml", utils.GetApplicationEnv())
	viper.SetConfigFile(configFileName)
	err := viper.ReadInConfig()
	if err != nil{
		return err
	}
	err = viper.Unmarshal(confObj)
	if err != nil {
		return err
	}
	return err
}